package jwt

type Issuer interface {
    GetID() uint
    GetEmail() string
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