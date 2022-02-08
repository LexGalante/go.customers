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

//GetAddresses -> get customers addresses
func GetAddresses(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idCustomer, err := strconv.ParseUint(params["id_customer"], 10, 64)
	if err != nil {
		BadRequest(w, r, models.MakeInvalidParameterError("id"))
		return
	}

	var addresses []entities.Address

	db := infrastructures.GetDatabaseConnection()
	if result := db.Where("customer_id = ?", idCustomer).Find(&addresses); result.Error != nil {
		InternalServerError(w, r, models.MakeUnexpectedError())
		return
	}

	if len(addresses) == 0 {
		NoContent(w, r)
		return
	}

	Ok(w, r, addresses)
}

//CreateAddress -> add new address into a customer
func CreateAddress(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idCustomer, err := strconv.ParseUint(params["id_customer"], 10, 64)
	if err != nil {
		BadRequest(w, r, models.MakeInvalidParameterError("id_customer"))
		return
	}

	var address entities.Address
	err = json.NewDecoder(r.Body).Decode(&address)
	if err != nil {
		InternalServerError(w, r, models.MakeInvalidJSONBodyError())
		return
	}

	address.CustomerID = idCustomer

	if address.PostalCode != "" {
		viaCepAddress, err := infrastructures.GetAddresFromViaCep(address.PostalCode)
		if err != nil {
			InternalServerError(w, r, models.MakeUnexpectedWithBodyError("cannot retrieve address from via cep"))
		} else {
			address.StreetName = viaCepAddress.StreetName
			address.District = viaCepAddress.District
			address.City = viaCepAddress.City
		}
	}

	Create(w, r, &address)
}

//DeleteAddress -> change new address into a customer
func DeleteAddress(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idAddress, err := strconv.ParseUint(params["id_address"], 10, 64)
	if err != nil {
		BadRequest(w, r, models.MakeInvalidParameterError("id_address"))
		return
	}

	Delete(w, r, &entities.Address{}, int(idAddress))
}
