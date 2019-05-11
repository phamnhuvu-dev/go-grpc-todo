package auth

import (
	"testing"
	"time"
)

func TestJWT(t *testing.T) {
	rawJWT, err := CreateJWT("1", "2", "3", "4", time.Minute*5)
	if err != nil {
		t.Error(err)
	}

	_, err = VerifyJWT(rawJWT)
	if err != nil {
		t.Error(err)
	}
}

func TestJWS(t *testing.T) {
	rawJWS, err := CreateJWS(JwsPayload{"Test", time.Now()})
	if err != nil {
		t.Error(err)
	}

	_, err = VerifyJWS(rawJWS)
	if err != nil {
		t.Error(err)
	}
}
