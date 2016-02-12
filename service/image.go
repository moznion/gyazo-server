package service

import (
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"

	"github.com/moznion/gyazo-server/aws"
)

func UploadImageForApp(fi multipart.File, s3 *aws.S3Info) (string, error) {
	fo, err := ioutil.TempFile("", "gyazo")
	if err != nil {
		return "", err
	}
	defer fo.Close()
	defer os.Remove(fo.Name())

	err = io.Copy(fo, fi)
	if err != nil {
		return "", err
	}

	return s3.Upload(fo, calcChecksum(fo)+".png", "image/png")
}
