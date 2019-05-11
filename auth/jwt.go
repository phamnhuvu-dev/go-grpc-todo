package auth

import (
	"gopkg.in/square/go-jose.v2"
	"gopkg.in/square/go-jose.v2/jwt"
	"time"
)

var jwtKey = []byte("secret")
var jwtSigner = createSignerJWT()

func createSignerJWT() jose.Signer {
	signer, err := jose.NewSigner(
		jose.SigningKey{Algorithm: jose.HS256, Key: jwtKey},
		(&jose.SignerOptions{}).WithType("JWT"),
	)
	if err != nil {
		panic(err)
	}
	return signer
}

func CreateJWT(
	id string, sub string,
	aud string, iss string,
	exp time.Duration) (string, error) {

	pubClaims := jwt.Claims{
		ID:       id,
		Subject:  sub,
		Audience: jwt.Audience{aud},
		Issuer:   iss,
		IssuedAt: jwt.NewNumericDate(time.Now()),
		Expiry:   jwt.NewNumericDate(time.Now().Add(exp)),
	}

	return jwt.Signed(jwtSigner).Claims(pubClaims).CompactSerialize()
}

func VerifyJWT(rawJWT string) (bool, error) {
	parsedJWT, err := jwt.ParseSigned(rawJWT)
	if err != nil {
		return false, err
	}

	out := jwt.Claims{}
	if err := parsedJWT.Claims(jwtKey, &out); err != nil {
		return false, err
	}

	return true, nil
}
