package utils

import (
	"fmt"
	"math/rand"
	"time"

	zjwt "github.com/zsmartex/pkg/jwt"
)

func RandomUID() string {
	rand.Seed(time.Now().UnixNano())
	min := 1000000000
	max := 9999999999
	return fmt.Sprintf("UID%v", rand.Intn(max-min+1)+min)
}

func ParserJWT(ss string) (*zjwt.Auth, error) {
	nks := zjwt.KeyStore{}
	nks.LoadPublicKeyFromFile("public.pem")

	tt, err := zjwt.ParseAndValidate(ss, nks.PublicKey)

	if err != nil {
		return nil, err
	}

	return &tt, err
}
