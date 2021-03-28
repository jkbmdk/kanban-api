package jwt

import "os"

type header struct {
	Algorithm string `json:"alg"`
	Type string `json:"typ"`
}

func createHeader() header {
	return header{Algorithm: os.Getenv("JWT_ALG"), Type: "JWT"}
}