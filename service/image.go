package service

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
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

	fo, err := ioutil.TempFile("", "gyazo")
	if err != nil {
		http.Error(w, "Internal server error", 500)
		return
	}
	defer fo.Close()
	defer os.Remove(fo.Name())

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

	fmt.Fprintf(w, "Succeeded")
}
