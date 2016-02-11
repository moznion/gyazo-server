package service

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

func PostImageFromClient(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Invalid request method", 405)
		return
	}

	fi, _, err := r.FormFile("imagedata")
	if err != nil {
		http.Error(w, "Cannot load imagedata", 400)
		return
	}
	defer fi.Close()

	p, err := generateTempFileName()
	if err != nil {
		http.Error(w, "Internal server error", 500)
		return
	}

	fo, err := os.Create(p)
	if err != nil {
		http.Error(w, "Internal server error", 500)
		return
	}
	defer os.Remove(p)

	buf := make([]byte, 1024)
	for {
		n, err := fi.Read(buf)
		if err != nil && err != io.EOF {
			http.Error(w, "Internal server error", 500)
			return
		}

		if n == 0 {
			break
		}

		_, err = fo.Write(buf[:n])
		if err != nil {
			http.Error(w, "Internal server error", 500)
			return
		}
	}
	fo.Close()

	fmt.Fprintf(w, "Succeeded")
}

func generateTempFileName() (string, error) {
	letterRunes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	rand.Seed(time.Now().UnixNano())

	cnt := 0
	var p string
	for {
		b := make([]rune, 128)
		for i := range b {
			b[i] = letterRunes[rand.Intn(len(letterRunes))]
		}

		p = strings.Join([]string{"/tmp/", string(b)}, "")
		_, err := os.Stat(p)
		if os.IsNotExist(err) {
			break
		}

		cnt++
		if cnt >= 10000 {
			return "", fmt.Errorf("Failed to generate temporary file name")
		}
	}

	return p, nil
}
