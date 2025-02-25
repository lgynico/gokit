package cryptox

import (
	"encoding/base64"
	"encoding/hex"
)

type Encoder func([]byte) string

var (
	HexEncoder          Encoder = hex.EncodeToString
	Base64URLEncoder    Encoder = base64.URLEncoding.EncodeToString
	Base64StdEncoder    Encoder = base64.StdEncoding.EncodeToString
	Base64RawURLEncoder Encoder = base64.RawURLEncoding.EncodeToString
	Base64RawStdEncoder Encoder = base64.RawStdEncoding.EncodeToString
)
