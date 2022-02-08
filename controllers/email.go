package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/lexgalante/go.customers/entities"
	"github.com/lexgalante/go.customers/infrastructures"
	"github.com/lexgalante/go.customers/models"
)

//GetEmails -> get customers emails
func GetEmails(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idCustomer, err := strconv.ParseUint(params["id_customer"], 10, 64)
	if err != nil {
		BadRequest(w, r, models.MakeInvalidParameterError("id"))
		return
	}

	var emails []entities.Email

	db := infrastructures.GetDatabaseConnection()
	if result := db.Where("customer_id = ?", idCustomer).Find(&emails); result.Error != nil {
		InternalServerError(w, r, models.MakeUnexpectedError())
		return
	}

	if len(emails) == 0 {
		NoContent(w, r)
		return
	}

	Ok(w, r, emails)
}

//CreateEmail -> add new address into a customer
func CreateEmail(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idCustomer, err := strconv.ParseUint(params["id_customer"], 10, 64)
	if err != nil {
		BadRequest(w, r, models.MakeInvalidParameterError("id_customer"))
		return
	}

	var email entities.Email
	err = json.NewDecoder(r.Body).Decode(&email)
	if err != nil {
		InternalServerError(w, r, models.MakeInvalidJSONBodyError())
		return
	}

	email.CustomerID = idCustomer

	Create(w, r, &email)
}

//DeleteEmail -> change new address into a customer
func DeleteEmail(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idEmail, err := strconv.ParseUint(params["id_email"], 10, 64)
	if err != nil {
		BadRequest(w, r, models.MakeInvalidParameterError("id_email"))
		return
	}

	Delete(w, r, &entities.Email{}, int(idEmail))
}
