package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/moznion/gyazo-server/service"
)

func (c *Controller) PostImageFromClient(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Invalid request method", 405)
		return
	}

	fi, _, err := r.FormFile("imagedata")
	if err != nil {
		c.renderErrorResponse(w, "Cannot load imagedata", 400)
		return
	}
	defer fi.Close()

	url, err := service.UploadImageForApp(fi, c.S3)
	if err != nil {
		c.renderErrorResponse(w, "Internal server error", 500)
		return
	}

	c.renderUploadedResponse(w, url)
}

func (c *Controller) renderUploadedResponse(w http.ResponseWriter, url string) {
	m := map[string]string{"message": "Secceeded", "url": url}
	s, _ := json.Marshal(m)
	fmt.Fprintf(w, string(s))
}
