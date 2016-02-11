package controller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/moznion/gyazo-server/service"
)

func (c *Controller) PostImageFromClient(w http.ResponseWriter, r *http.Request) {
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

	url, err := service.UploadImageForApp(fi, c.S3)
	if err != nil {
		http.Error(w, "Internal server error", 500)
		return
	}

	ps := strings.Split(url, "/")

	fmt.Fprintf(w, strings.Join([]string{c.Host, ps[len(ps)-1]}, "/"))
}

func (c *Controller) GetImage(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Invalid request method", 405)
		return
	}

	vs := mux.Vars(r)
	key := vs["key"]

	o, err := c.S3.Get(key)
	if err != nil {
		http.Error(w, "Not found", 404)
		return
	}

	w.Header().Set("Content-Type", *o.ContentType)
	w.WriteHeader(http.StatusOK)

	b, _ := ioutil.ReadAll(o.Body)
	w.Write(b)
}
