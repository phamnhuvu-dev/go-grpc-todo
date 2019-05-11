package auth

import (
	"crypto/rand"
	"crypto/rsa"
	"gopkg.in/square/go-jose.v2"
	"time"
)

type JwsPayload struct {
	Name      string
	CreatedAt time.Time
}

func CreatePrivateKey() *rsa.PrivateKey {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}
	return privateKey
}

func CreatePrivateKeySigner(privateKey *rsa.PrivateKey) jose.Signer {
	signer, err := jose.NewSigner(jose.SigningKey{Algorithm: jose.PS512, Key: privateKey}, nil)
	if err != nil {
		panic(err)
	}
	return signer
}


