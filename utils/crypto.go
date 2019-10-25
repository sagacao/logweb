package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func md5sum(str []byte) string {
	l := md5.New()
	l.Write(str)
	return hex.EncodeToString(l.Sum(nil))
}
