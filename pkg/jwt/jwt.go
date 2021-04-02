package jwt

import (
	"encoding/json"
	"errors"
	"strings"
)

func Generate(issuer Issuer) string {
	header := createHeader()
	headerJSON, _ := json.Marshal(header)
	headerEncoded := base64Encode(headerJSON)

	payload := createPayload(issuer)
	payloadJSON, _ := json.Marshal(payload)
	payloadEncoded := base64Encode(payloadJSON)

	return strings.Join([]string{headerEncoded, payloadEncoded, signature(headerEncoded, payloadEncoded)}, ".")
}

func Parse(token string) (uint, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return 0, errors.New("JWT VERIFY: wrong token format")
	}

	if signature(parts[0], parts[1]) != parts[2] {
		return 0, errors.New("JWT VERIFY: sum mismatch")
	}

	payloadJSON, err := base64Decode(parts[1])
	if err != nil {
		return 0, errors.New("JWT VERIFY: unable to read payloadJSON")
	}

	var payload payload
	if json.Unmarshal(payloadJSON, &payload) != nil {
		return 0, errors.New("JWT VERIFY: payload properties mismatch")
	}

	if payload.expired() {
		return 0, errors.New("JWT VERIFY: token expired")
	}

	return payload.Issuer, nil
}
