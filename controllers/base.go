package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"runtime"

	"github.com/lexgalante/go.customers/models"
)

//Ok -> 200
func Ok(w http.ResponseWriter, r *http.Request, v interface{}) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(v)
}

//Created -> 201
func Created(w http.ResponseWriter, r *http.Request, v interface{}) {
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(v)
}

//Accepted -> 202
func Accepted(w http.ResponseWriter, r *http.Request, v interface{}) {
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(v)
}

//NoContent -> 204
func NoContent(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}

//BadRequest -> 400
func BadRequest(w http.ResponseWriter, r *http.Request, m models.Error) {
	pc, _, _, ok := runtime.Caller(1)
	details := runtime.FuncForPC(pc)
	if ok && details != nil {
		log.Println(fmt.Sprintf("[%s] - raise a bad request: %s", details.Name(), m))
	}

	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(m)
}

//NotFound -> 404
func NotFound(w http.ResponseWriter, r *http.Request, m models.Error) {
	pc, _, _, ok := runtime.Caller(1)
	details := runtime.FuncForPC(pc)
	if ok && details != nil {
		log.Println(fmt.Sprintf("[%s] - raise a not found: %s", details.Name(), m))
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(m)
}

//UnprocessableEntity -> 422
func UnprocessableEntity(w http.ResponseWriter, r *http.Request, v interface{}) {
	pc, _, _, ok := runtime.Caller(1)
	details := runtime.FuncForPC(pc)
	if ok && details != nil {
		log.Println(fmt.Sprintf("[%s] - raise a unprocessable entity: %s", details.Name(), v))
	}

	w.WriteHeader(http.StatusUnprocessableEntity)
	json.NewEncoder(w).Encode(v)
}

//InternalServerError -> 500
func InternalServerError(w http.ResponseWriter, r *http.Request, m models.Error) {
	pc, _, _, ok := runtime.Caller(1)
	details := runtime.FuncForPC(pc)
	if ok && details != nil {
		log.Println(fmt.Sprintf("[%s] - raise a internal server error: %s", details.Name(), m))
	}

	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(m)
}
