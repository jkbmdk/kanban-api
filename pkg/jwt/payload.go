package jwt

import (
    "os"
    "strconv"
    "time"
)

type payload struct {
    Email    string `json:"email"`
    Issuer   uint   `json:"iss"`
    IssuedAt int64  `json:"iat"`
    ExpireIn uint   `json:"exp"`
}

func (p *payload) expired() bool {
    return time.Now().Unix() > p.IssuedAt+int64(p.ExpireIn)
}

func createPayload(issuer Issuer) payload {
    e, _ := strconv.ParseUint(os.Getenv("JWT_EXP"), 10, 32)
    at := time.Now().Unix()
    return payload{
        Email:    issuer.GetEmail(),
        Issuer:   issuer.GetID(),
        IssuedAt: at,
        ExpireIn: uint(e),
    }
}
