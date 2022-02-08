package infrastructures

import "github.com/golang-jwt/jwt"

//UserCustomClaims -> custom claims for golang-jwt
type UserCustomClaims struct {
	*jwt.StandardClaims
	IsAdmin bool   `json:"is_admin"`
	Email   string `json:"email"`
}
