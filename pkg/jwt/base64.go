package jwt

import "encoding/base64"

func base64Encode(message []byte) string {
    return base64.URLEncoding.EncodeToString(message)
}

func base64Decode(encoded string) ([]byte, error) {
    return base64.URLEncoding.DecodeString(encoded)
}
