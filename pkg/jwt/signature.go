package jwt

import (
    "crypto/hmac"
    "crypto/sha512"
)

func signature(header string, payload string, secret string) string {
    h := hmac.New(sha512.New, []byte(secret))
    h.Write([]byte(header + "." + payload))
    return base64Encode(h.Sum(nil))
}
