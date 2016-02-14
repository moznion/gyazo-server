package service

import (
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"

	"github.com/moznion/gyazo-server/aws"
)

func UploadImage(fi multipart.File, s3 *aws.S3Info) (string, error) {
	fo, err := ioutil.TempFile("", "gyazo")
	if err != nil {
		return "", err
	}
	defer fo.Close()
	defer os.Remove(fo.Name())

	_, err = io.Copy(fo, fi)
	if err != nil {
		return "", err
	}

	contentType, ex, err := judgeContentType(fi)
	if err != nil {
		return "", err
	}

	return s3.Upload(fo, calcChecksum(fo)+ex, contentType)
}

func judgeContentType(f multipart.File) (string, string, error) {
	f.Seek(0, 0)
	buf := make([]byte, 16)
	count, err := f.Read(buf)
	if err != nil {
		return "", "", err
	}

	if count >= 10 && string(buf[6:10]) == "JFIF" {
		return "image/jpeg", ".jpg", nil
	}
	if count >= 4 && string(buf[0:3]) == "GIF" {
		return "image/gif", ".gif", nil
	}
	if count >= 2 && string(buf[1:4]) == "PNG" {
		return "image/png", ".png", nil
	}

	return "application/octet-stream", "", nil
}
