package jwt

import (
	"encoding/json"
	"os"
	"strings"
	"testing"
	"time"
)

var i = issuer{Email: "jkbmdk@example.com", ID: 0}

func TestGenerateProvidesExpectedHeader(t *testing.T) {
	expected := header{Algorithm: "HS512", Type: "JWT"}

	_ = os.Setenv("JWT_ALG", "HS512")
	token := Generate(&i)
	parts := strings.Split(token, ".")

	j, _ := base64Decode(parts[0])
	var result header
	if json.Unmarshal(j, &result) != nil {
		t.Error("unable to decode header")
	}

	if result != expected {
		t.Errorf("result header %v mismatch %v", result, expected)
	}
}

func TestGenerateProvidesExpectedPayload(t *testing.T) {
	token := Generate(&i)
	parts := strings.Split(token, ".")

	j, _ := base64Decode(parts[1])
	var result payload
	if json.Unmarshal(j, &result) != nil {
		t.Error("unable to decode payload")
	}

	if result.Issuer != i.GetID() {
		t.Errorf("payload issuer ID mismatch result: %v, expected: %v", result.Issuer, i.GetID())
	}

	if result.Email != i.GetEmail() {
		t.Errorf("payload issuer ID mismatch result: %v, expected: %v", result.Email, i.GetEmail())
	}
}

func TestVerifyAcceptCorrectToken(t *testing.T) {
	_ = os.Setenv("JWT_ALG", "HS512")
	_ = os.Setenv("JWT_SECRET", "secret")
	
	token := Generate(&i)

	if err := Verify(token); err != nil {
		t.Errorf("token verifies fail")
	}
}

func TestVerifyDoesAcceptTokenUsingDifferentSecret(t *testing.T) {
	_ = os.Setenv("JWT_ALG", "HS512")
	_ = os.Setenv("JWT_SECRET", "secret")

	token := Generate(&i)

	_ = os.Setenv("JWT_SECRET", "newsecret")

	if err := Verify(token); err == nil {
		t.Errorf("token verifier shall not accept token hashed with another secret")
	}
}

func TestVerifyDoesAcceptExpiredToken(t *testing.T) {
	_ = os.Setenv("JWT_ALG", "HS512")
	_ = os.Setenv("JWT_SECRET", "secret")
	_ = os.Setenv("JWT_EXP", "0")

	token := Generate(&i)

	time.Sleep(time.Second)

	if err := Verify(token); err == nil {
		t.Errorf("token verifier shall not accept expired token")
	}
}