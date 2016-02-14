package controller

import (
	"net/http"

	"github.com/moznion/gyazo-server/aws"
)

type Controller struct {
	S3         *aws.S3Info
	Host       string
	Passphrase string
}

func NewController(s3 *aws.S3Info, host, passphrase string) *Controller {
	return &Controller{
		S3:         s3,
		Host:       host,
		Passphrase: passphrase,
	}
}

func (c *Controller) authenticate(r *http.Request) bool {
	if c.Passphrase == "" || c.Passphrase == r.Header.Get("X-Gyazo-Auth") {
		return true
	}
	return false
}

func (c *Controller) isPost(r *http.Request) bool {
	if r.Method == "POST" {
		return true
	}
	return false
}
