package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/lexgalante/go.customers/src/entities"
	"github.com/lexgalante/go.customers/src/infrastructures"
	"github.com/lexgalante/go.customers/src/models"
	"gorm.io/gorm"
)

//Register -> register new user
func Register(w http.ResponseWriter, r *http.Request) {
	var user entities.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		InternalServerError(w, r, models.MakeInvalidJSONBodyError())
		return
	}

	db := infrastructures.GetDatabaseConnection()

	validations, err := user.Validate()

	var userExist entities.Customer
	db.Select("email").Where("email = ?", user.Email).Limit(1).Find(&userExist)
	if userExist.Login != "" {
		validations["EMAIL_ALREADY_EXIST"] = fmt.Sprintf("email %s already exists", userExist.Login)
	}

	if err != nil {
		InternalServerError(w, r, models.MakeUnexpectedError())
		return
	}

	user.IsAdmin = false
	user.CryptPassword()

	result := db.Create(&user)
	if result.Error != nil {
		InternalServerError(w, r, models.MakeUnexpectedWithBodyError(result.Error.Error()))
		return
	}

	Created(w, r, user)
}

//Login -> get jwk access token
func Login(w http.ResponseWriter, r *http.Request) {
	var user entities.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		InternalServerError(w, r, models.MakeInvalidJSONBodyError())
		return
	}

	db := infrastructures.GetDatabaseConnection()
	var userInDb entities.User
	if result := db.First(&userInDb, "email = ?", user.Email); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			NotFound(w, r, models.MakeNotFoundError("user"))
			return
		}
		InternalServerError(w, r, models.MakeUnexpectedError())
		return
	}

	if !userInDb.VerifyPassword(user.Password) {
		Unauthorized(w, r, models.MakeUnauthorizedError())
		return
	}

	accessToken, expiresAt, err := userInDb.CreateToken()
	if err != nil {
		InternalServerError(w, r, models.MakeUnexpectedWithBodyError(err.Error()))
		return
	}

	Ok(w, r, models.AuthToken{TokenType: "Bearer ", AccessToken: accessToken, ExpiresAt: expiresAt})
}
