package controller

import (
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
