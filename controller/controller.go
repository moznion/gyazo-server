package controller

import (
	"encoding/json"
	"net/http"

	"github.com/moznion/gyazo-server/aws"
)

type Controller struct {
	S3 *aws.S3Info
}

func NewController(s3 *aws.S3Info) *Controller {
	return &Controller{
		S3: s3,
	}
}

func (c *Controller) renderErrorResponse(w http.ResponseWriter, errMsg string, code int) {
	m := map[string]string{"message": errMsg}
	s, _ := json.Marshal(m)
	http.Error(w, string(s), code)
}
