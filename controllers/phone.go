package controllers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/lexgalante/go.customers/entities"
	"github.com/lexgalante/go.customers/infrastructures"
	"github.com/lexgalante/go.customers/models"
)

//GetPhones -> get customers phones
func GetPhones(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		BadRequest(w, r, models.MakeInvalidParameterError("id"))
		return
	}

	var phones []entities.Phone

	db := infrastructures.GetDatabaseConnection()
	if result := db.Where("customer_id = ?", id).Find(&phones); result.Error != nil {
		InternalServerError(w, r, models.MakeUnexpectedError())
		return
	}

	if len(phones) == 0 {
		NoContent(w, r)
		return
	}

	Ok(w, r, phones)
}
