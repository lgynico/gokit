package cryptox

import (
	"crypto/md5"
)

func MD5(source string) []byte {
	hash := md5.New()
	_, _ = hash.Write([]byte(source))
	return hash.Sum(nil)
}

func MD5String(source string, encode Encoder) string {
	return encode(MD5(source))
}
