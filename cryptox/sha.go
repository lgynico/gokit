package cryptox

import (
	"crypto/hmac"
	"crypto/sha256"
)

func SHA256(key []byte, source string) []byte {
	hash := hmac.New(sha256.New, key)
	_, _ = hash.Write([]byte(source))
	return hash.Sum(nil)
}

func SHA256String(key []byte, source string, encode Encoder) string {
	return encode(SHA256(key, source))
}
