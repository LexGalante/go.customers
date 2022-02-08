package entities

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/lexgalante/go.customers/infrastructures"
	"github.com/lexgalante/go.customers/utils"
	"golang.org/x/crypto/bcrypt"
)

//User -> represent an user
type User struct {
	Email    string `json:"email" gorm:"primaryKey"`
	Password string `json:"password" gorm:"type:varchar(200);not null"`
	IsAdmin  bool   `json:"is_admin" gorm:"default:false"`
}

//CryptPassword -> crypt password
func (u *User) CryptPassword() error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
	if err != nil {
		return err
	}

	u.Password = string(bytes)

	return nil
}

//VerifyPassword -> check password
func (u *User) VerifyPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))

	return err == nil
}

//CreateToken -> return jwk user token
func (u *User) CreateToken() (string, int64, error) {
	expireAt := time.Now().Add(time.Hour * 1).Unix()
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, infrastructures.UserCustomClaims{
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: expireAt,
		},
		IsAdmin: u.IsAdmin,
		Email:   u.Email,
	})

	token, err := claims.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", expireAt, err
	}

	return token, expireAt, nil
}

//Validate -> validate instance
func (u *User) Validate() (map[string]string, error) {
	errors := make(map[string]string)

	if u.Email == "" || len(u.Email) > 250 || !utils.IsEmailValid(u.Email) {
		errors["INVALID_EMAIL"] = "is mandatory and must contain a maximum of 250 characters and valid email"
	}

	if !utils.IsValidPassword(u.Password) {
		errors["INVALID_PASSWORD"] = "is mandatory and must contain number a capitalize char, a number and minimum length 7"
	}

	return errors, nil
}
