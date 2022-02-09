package main

import (
	"log"
	"net/http"

	"github.com/lexgalante/go.customers/src/controllers"
	"github.com/lexgalante/go.customers/src/middlewares"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("unable to load .env file")
	}

	router := mux.NewRouter()
	addHandlers(router)
	addMiddlewares(router)

	log.Fatal(http.ListenAndServe(":5000", router))
}

func addHandlers(router *mux.Router) {
	router.HandleFunc("/register", controllers.Register).Methods(http.MethodPost)
	router.HandleFunc("/login", controllers.Login).Methods(http.MethodPost)
	subrouter := router.PathPrefix("/v1/").Subrouter()
	subrouter.HandleFunc("/customers", controllers.GetCustomers).
		Queries("page", "{page:[0-9,]+}", "page_size", "{page_size:[0-9,]+}").
		Schemes("http", "https").
		Methods(http.MethodGet)
	subrouter.HandleFunc("/customers", controllers.CreateCustomer).Methods(http.MethodPost)
	subrouter.HandleFunc("/customers/{id_customer:[0-9,]+}", controllers.GetCustomerByID).Methods(http.MethodGet)
	subrouter.HandleFunc("/customers/{id_customer:[0-9,]+}", controllers.UpdateCustomer).Methods(http.MethodPut)
	subrouter.HandleFunc("/customers/{id_customer:[0-9,]+}", controllers.DeleteCustomer).Methods(http.MethodDelete)
	subrouter.HandleFunc("/customers/{id_customer:[0-9,]+}/addresses", controllers.GetAddresses).Methods(http.MethodGet)
	subrouter.HandleFunc("/customers/{id_customer:[0-9,]+}/addresses", controllers.CreateAddress).Methods(http.MethodPost)
	subrouter.HandleFunc("/customers/{id_customer:[0-9,]+}/addresses/{id_address:[0-9,]+}", controllers.DeleteAddress).Methods(http.MethodDelete)
	subrouter.HandleFunc("/customers/{id_customer:[0-9,]+}/emails", controllers.GetEmails).Methods(http.MethodGet)
	subrouter.HandleFunc("/customers/{id_customer:[0-9,]+}/emails", controllers.CreateEmail).Methods(http.MethodPost)
	subrouter.HandleFunc("/customers/{id_customer:[0-9,]+}/emails/{id_email:[0-9,]+}", controllers.DeleteEmail).Methods(http.MethodDelete)
	subrouter.HandleFunc("/customers/{id_customer:[0-9,]+}/phones", controllers.GetPhones).Methods(http.MethodGet)
	subrouter.HandleFunc("/customers/{id_customer:[0-9,]+}/phones", controllers.CreatePhone).Methods(http.MethodPost)
	subrouter.HandleFunc("/customers/{id_customer:[0-9,]+}/phones/{id_phone:[0-9,]+}", controllers.DeletePhone).Methods(http.MethodDelete)
}

func addMiddlewares(router *mux.Router) {
	router.Use(middlewares.JwtSecurityMiddleware)
	router.Use(mux.CORSMethodMiddleware(router))
}
