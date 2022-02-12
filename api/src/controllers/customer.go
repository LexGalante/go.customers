package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/lexgalante/go.customers/api/src/entities"
	"github.com/lexgalante/go.customers/api/src/infrastructures"
	"github.com/lexgalante/go.customers/api/src/models"
	"gorm.io/gorm"

	"github.com/gorilla/mux"
)

//GetCustomers -> paginate customers
func GetCustomers(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	page, err := strconv.ParseInt(params["page"], 10, 8)
	if err != nil {
		BadRequest(w, r, models.MakeInvalidParameterError("page"))
		return
	}

	pageSize, err := strconv.ParseInt(params["page_size"], 10, 8)
	if err != nil {
		BadRequest(w, r, models.MakeInvalidParameterError("page_size"))
		return
	}

	limit := pageSize
	offset := (page - 1) * limit

	var customers []entities.Customer

	db := infrastructures.GetDatabaseConnection()
	if result := db.Where("active = true").Limit(int(limit)).Offset(int(offset)).Find(&customers); result.Error != nil {
		InternalServerError(w, r, models.MakeUnexpectedError())
		return
	}

	if len(customers) == 0 {
		NoContent(w, r)
		return
	}

	Ok(w, r, customers)
}

//GetCustomerByID -> retrieve a single customer by id
func GetCustomerByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idCustomer, err := strconv.ParseUint(params["id_customer"], 10, 64)
	if err != nil {
		BadRequest(w, r, models.MakeInvalidParameterError("id_customer"))
		return
	}

	customer := entities.Customer{}
	db := infrastructures.GetDatabaseConnection()
	if result := db.Preload("Addresses").Preload("Emails").Preload("Phones").First(&customer, idCustomer); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			NotFound(w, r, models.MakeNotFoundError("customer"))
			return
		}
		InternalServerError(w, r, models.MakeUnexpectedError())
		return
	}

	Ok(w, r, customer)
}

//CreateCustomer -> create new customer
func CreateCustomer(w http.ResponseWriter, r *http.Request) {
	var customer entities.Customer
	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		InternalServerError(w, r, models.MakeInvalidJSONBodyError())
		return
	}

	db := infrastructures.GetDatabaseConnection()

	validations, err := customer.Validate()
	if err != nil {
		InternalServerError(w, r, models.MakeUnexpectedError())
		return
	}

	var customerExist entities.Customer
	db.Select("login").Where("login = ?", customer.Login).Limit(1).Find(&customerExist)
	if customerExist.Login != "" {
		validations["LOGIN_ALREADY_EXIST"] = fmt.Sprintf("login %s already exists", customerExist.Login)
	}

	if len(validations) > 0 {
		UnprocessableEntity(w, r, validations)
		return
	}

	customer.CreatedAt = time.Now()
	customer.Active = true

	result := db.Create(&customer)
	if result.Error != nil {
		InternalServerError(w, r, models.MakeUnexpectedWithBodyError(result.Error.Error()))
		return
	}

	Created(w, r, customer)

	go sendMessage("new_customers", customer)
}

//UpdateCustomer -> apply changes at single customer
func UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idCustomer, err := strconv.ParseUint(params["id_customer"], 10, 64)
	if err != nil {
		BadRequest(w, r, models.MakeInvalidParameterError("id_customer"))
		return
	}

	var customer entities.Customer
	err = json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		InternalServerError(w, r, models.MakeInvalidJSONBodyError())
		return
	}

	db := infrastructures.GetDatabaseConnection()

	if customer.ID == 0 {
		customer.ID = idCustomer
	}

	validations, err := customer.Validate()
	if err != nil {
		InternalServerError(w, r, models.MakeUnexpectedError())
		return
	}
	if len(validations) > 0 {
		UnprocessableEntity(w, r, validations)
		return
	}

	dbCustomer := entities.Customer{}
	if result := db.First(&dbCustomer, idCustomer); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			NotFound(w, r, models.MakeNotFoundError("customer"))
			return
		}
		InternalServerError(w, r, models.MakeUnexpectedError())
		return
	}

	dbCustomer.FirstName = customer.FirstName
	dbCustomer.LastName = customer.LastName
	dbCustomer.Active = customer.Active
	dbCustomer.UpdatedAt = time.Now()

	result := db.Save(&dbCustomer)
	if result.Error != nil {
		InternalServerError(w, r, models.MakeUnexpectedWithBodyError(result.Error.Error()))
		return
	}

	Accepted(w, r, customer)

	go sendMessage("changes_customers", customer)
}

//DeleteCustomer -> delete customer by id
func DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idCustomer, err := strconv.ParseUint(params["id_customer"], 10, 64)
	if err != nil {
		BadRequest(w, r, models.MakeInvalidParameterError("id_customer"))
		return
	}

	db := infrastructures.GetDatabaseConnection()
	if result := db.Delete(&entities.Customer{}, idCustomer); result.Error != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			NotFound(w, r, models.MakeNotFoundError("customer"))
			return
		}
		InternalServerError(w, r, models.MakeUnexpectedError())
		return
	}

	NoContent(w, r)
}

func sendMessage(exchangeName string, customer entities.Customer) {
	message, err := json.Marshal(customer)
	if err != nil {
		log.Fatalf("unexpected error ocurred while trying to serialize entities.Customer: %s", err.Error())
		return
	}

	err = infrastructures.PublishMessage(exchangeName, message)
	if err != nil {
		log.Fatalf("unexpected error ocurred while trying to publish message: %s, in exchange: %s", exchangeName, message)
		return
	}
}
