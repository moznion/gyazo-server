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

	buf := make([]byte, 1024)
	for {
		n, err := fi.Read(buf)
		if err != nil && err != io.EOF {
			return "", err
		}

		if n == 0 {
			break
		}

		_, err = fo.Write(buf[:n])
		if err != nil {
			return "", err
		}
	}

	return s3.Upload(fo, calcChecksum(fo)+".png", "image/png")
}
