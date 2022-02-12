package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/lexgalante/go.customers/api/src/entities"
	"github.com/lexgalante/go.customers/api/src/infrastructures"
	"github.com/lexgalante/go.customers/api/src/models"
)

//GetPhones -> get customers phones
func GetPhones(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idCustomer, err := strconv.ParseUint(params["id_customer"], 10, 64)
	if err != nil {
		BadRequest(w, r, models.MakeInvalidParameterError("id"))
		return
	}

	var phones []entities.Phone

	db := infrastructures.GetDatabaseConnection()
	if result := db.Where("customer_id = ?", idCustomer).Find(&phones); result.Error != nil {
		InternalServerError(w, r, models.MakeUnexpectedError())
		return
	}

	if len(phones) == 0 {
		NoContent(w, r)
		return
	}

	Ok(w, r, phones)
}

//CreatePhone -> add new address into a customer
func CreatePhone(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idCustomer, err := strconv.ParseUint(params["id_customer"], 10, 64)
	if err != nil {
		BadRequest(w, r, models.MakeInvalidParameterError("id_customer"))
		return
	}

	var phone entities.Phone
	err = json.NewDecoder(r.Body).Decode(&phone)
	if err != nil {
		InternalServerError(w, r, models.MakeInvalidJSONBodyError())
		return
	}

	phone.CustomerID = idCustomer

	Create(w, r, &phone)
}

//DeletePhone -> change new address into a customer
func DeletePhone(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idPhone, err := strconv.ParseUint(params["id_phone"], 10, 64)
	if err != nil {
		BadRequest(w, r, models.MakeInvalidParameterError("id_phone"))
		return
	}

	Delete(w, r, &entities.Phone{}, int(idPhone))
}
