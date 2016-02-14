package controller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/moznion/gyazo-server/service"
)

func (c *Controller) PostImage(w http.ResponseWriter, r *http.Request) {
	if !c.isPost(r) {
		http.Error(w, "Invalid request method", 405)
		return
	}

	if !c.authenticate(r) {
		http.Error(w, "Forbidden", 403)
		return
	}

	fi, _, err := r.FormFile("imagedata")
	if err != nil {
		http.Error(w, "Cannot load imagedata", 400)
		return
	}
	defer fi.Close()

	url, err := service.UploadImage(fi, c.S3)
	if err != nil {
		http.Error(w, "Internal server error", 500)
		return
	}

	ps := strings.Split(url, "/")

	fmt.Fprintf(w, strings.Join([]string{c.Host, ps[len(ps)-1]}, "/"))
}

func (c *Controller) GetImage(w http.ResponseWriter, r *http.Request) {
	if !c.isGet(r) {
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
