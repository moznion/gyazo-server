package controller

import (
	"github.com/moznion/gyazo-server/aws"
)

type Controller struct {
	S3   *aws.S3Info
	Host string
}

func NewController(s3 *aws.S3Info, host string) *Controller {
	return &Controller{
		S3:   s3,
		Host: host,
	}
}
