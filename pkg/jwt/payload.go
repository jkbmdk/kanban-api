package jwt

import (
    "os"
    "strconv"
    "time"
)

type Issuer interface {
    GetID() uint
    GetEmail() string
}

type payload struct {
    Email string `json:"email"`
    Iss   uint   `json:"iss"`
    Iat   int64  `json:"iat"`
    Exp   uint   `json:"exp"`
}

func (p *payload) expired() bool {
    return time.Now().Unix() > p.Iat+int64(p.Exp)
}

func createPayload(issuer Issuer) payload {
    e, _ := strconv.ParseUint(os.Getenv("JWT_EXP"), 2, 32)
    at := time.Now().Unix()
    return payload{
        Email: issuer.GetEmail(),
        Iss:   issuer.GetID(),
        Iat:   at,
        Exp:   uint(e),
    }
}

type issuer struct {
    ID    uint
    Email string
}

func (i *issuer) GetID() uint {
    return i.ID
}

func (i *issuer) GetEmail() string {
    return i.Email
}
