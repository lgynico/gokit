package jwtx

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

func Decode(signKey []byte, token string) (map[string]any, error) {
	// tk, err := jwt.Parse(token, func(t *jwt.Token) (any, error) {
	// 	return signKey, nil
	// })

	tk, err := jwt.ParseWithClaims(token, jwt.MapClaims{}, func(t *jwt.Token) (any, error) {
		return signKey, nil
	})

	if err != nil {
		return nil, err
	}

	if tk.Valid {
		return tk.Claims.(jwt.MapClaims), nil
	}

	return nil, errors.New("invalid token")
}
