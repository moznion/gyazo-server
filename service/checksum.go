package service

import (
	"crypto/sha256"
	"encoding/base64"
	"io"
	"math"
	"os"
)

const filechunk = 8192

func calcChecksum(f *os.File) string {
	info, _ := f.Stat()
	size := info.Size()

	blocks := uint64(math.Ceil(float64(size) / float64(filechunk)))

	hash := sha256.New()

	for i := uint64(0); i < blocks; i++ {
		blocksize := int(math.Min(filechunk, float64(size-int64(i*filechunk))))
		buf := make([]byte, blocksize)

		f.Read(buf)
		io.WriteString(hash, string(buf)) // append into the hash
	}

	return base64.URLEncoding.EncodeToString(hash.Sum(nil))
}
