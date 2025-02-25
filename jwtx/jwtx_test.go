package jwtx

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/bwoil/erase-game/app"
	"github.com/bwoil/erase-game/utils/cryptox"
)

func TestHS256(t *testing.T) {
	fmt.Println(HmacSha256([]byte("ILoveGames"), map[string]any{
		"username": "nico",
	}))

	var header = map[string]any{
		"alg": "HS256",
		"typ": "JWT",
	}
	var claim = map[string]any{
		"username": "nico",
	}

	h, _ := json.Marshal(header)
	c, _ := json.Marshal(claim)
	t1 := base64.RawURLEncoding.EncodeToString(h)
	t2 := base64.RawURLEncoding.EncodeToString(c)

	t3 := cryptox.SHA256String([]byte("ILoveGames"), t1+"."+t2, cryptox.Base64RawURLEncoder)
	fmt.Println(t1 + "." + t2 + "." + t3)

}

func TestHS256Decode(t *testing.T) {
	var claim = map[string]any{
		"username": "nico",
	}

	token, err := HmacSha256(app.GameContext.SignKey(), claim)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(token)
	// token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOjEsInVzZXJuYW1lIjoibmljbzAwMSJ9.I_YYDzp39dw6r2bMm8jagZmCXXttEnY3Ww8r8MOIeWQ"

	m, err := Decode(app.GameContext.SignKey(), token)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(m)
}
