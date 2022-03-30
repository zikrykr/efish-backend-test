package model

import "github.com/dgrijalva/jwt-go"

type TokenClaims struct {
	*jwt.StandardClaims
	Name      string `json:"name"`
	Phone     string `json:"phone"`
	Role      string `json:"role"`
	CreatedAt string `json:"createdAt"`
}
