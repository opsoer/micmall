package models

import (
	"github.com/dgrijalva/jwt-go"
)
// jwt

type CustomClaims struct {
	ID          uint
	NickName    string
	AuthorityId uint
	jwt.StandardClaims
}
