package jwt

import (
    "crypto/hmac"
    "crypto/sha512"
    "os"
)

func signature(header string, payload string) string {
    h := hmac.New(sha512.New, []byte(os.Getenv("JWT_SECRET")))
    h.Write([]byte(header + "." + payload))
    return base64Encode(h.Sum(nil))
}
