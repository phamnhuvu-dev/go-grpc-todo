package auth

import (
	"encoding/json"
	"gopkg.in/square/go-jose.v2"
)

var jwsPrivateKey = CreatePrivateKey()
var jwsSigner = CreatePrivateKeySigner(jwsPrivateKey)

func CreateJWS(object interface{}) (string, error) {
	bytes, err := json.Marshal(object)

	jws, err := jwsSigner.Sign(bytes)
	if err != nil {
		panic(err)
	}
	return jws.CompactSerialize()
}

func VerifyJWS(rawJWS string) ([]byte, error) {
	jws, err := jose.ParseSigned(rawJWS)
	if err != nil {
		return nil, err
	}

	output, err := jws.Verify(&jwsPrivateKey.PublicKey)
	if err != nil {
		return nil, err
	}
	return output, nil
}
