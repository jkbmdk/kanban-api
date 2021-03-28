package jwt

import (
    "encoding/json"
    "errors"
    "os"
    "strings"
)

var secret = os.Getenv("JWT_SECRET")

func Generate(issuer Issuer) string {
    header := createHeader()
    headerJSON, _ := json.Marshal(header)
    headerEncoded := base64Encode(headerJSON)

    payload := createPayload(issuer)
    payloadJSON, _ := json.Marshal(payload)
    payloadEncoded := base64Encode(payloadJSON)

    return strings.Join([]string{headerEncoded, payloadEncoded, signature(headerEncoded, payloadEncoded, secret)}, ".")
}

func Verify(token string) (Issuer, error) {
    parts := strings.Split(token, ".")
    if len(parts) != 3 {
        return nil, errors.New("JWT VERIFY: wrong token format")
    }

    payloadJSON, err := base64Decode(parts[1])
    if err != nil {
        return nil, errors.New("JWT VERIFY: unable to read payloadJSON")
    }
    var payload payload
    if json.Unmarshal(payloadJSON, &payload) != nil {
        return nil, errors.New("JWT VERIFY: payload properties mismatch")
    }

    if payload.expired() {
        return nil, errors.New("JWT VERIFY: token expired")
    }

    if signature(parts[0], parts[1], secret) != parts[2] {
        return nil, errors.New("JWT VERIFY: sum mismatch")
    }

    i := issuer{ID: payload.Iss, Email: payload.Email}
    return &i, nil
}