package jwt

import (
    "encoding/base64"
    "strings"
)

func base64Encode(message []byte) string {
    return strings.TrimRight(base64.URLEncoding.EncodeToString(message), "=")
}

func base64Decode(encoded string) ([]byte, error) {
    switch len(encoded) % 4 {
    case 2:
        encoded += "=="
    case 3:
        encoded += "="
    }
    return base64.URLEncoding.DecodeString(encoded)
}
