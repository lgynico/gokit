package jwtx

import "github.com/golang-jwt/jwt/v5"

func HmacSha256(signKey []byte, claims map[string]any) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims(claims))
	return token.SignedString(signKey)
}

func HmacSha384() {}

func HmacSha512() {}

func Rsa256() {}

func Rsa384() {}

func Rsa512() {}

func Ecdsa256() {}

func Ecdsa384() {}

func Ecdsa512() {}
