package main

import (
	"log"
	"net/http"

	"github.com/lexgalante/go.customers/controllers"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Unable to load .env file")
	}

	router := prepareRouterHandler()

	log.Fatal(http.ListenAndServe(":5000", router))
}

func prepareRouterHandler() *mux.Router {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/v1/").Subrouter()
	subrouter.HandleFunc("/customers", controllers.GetCustomers).
		Queries("page", "{page:[0-9,]+}", "page_size", "{page_size:[0-9,]+}").
		Schemes("http", "https").
		Methods(http.MethodGet)
	subrouter.HandleFunc("/customers", controllers.CreateCustomer).Methods(http.MethodPost)
	subrouter.HandleFunc("/customers/{id:[0-9,]+}", controllers.GetCustomerByID).Methods(http.MethodGet)
	subrouter.HandleFunc("/customers/{id:[0-9,]+}", controllers.UpdateCustomer).Methods(http.MethodPut)
	subrouter.HandleFunc("/customers/{id:[0-9,]+}", controllers.DeleteCustomer).Methods(http.MethodDelete)
	subrouter.HandleFunc("/customers/{id:[0-9,]+}/addresses", controllers.GetAddresses).Methods(http.MethodGet)
	subrouter.HandleFunc("/customers/{id:[0-9,]+}/emails", controllers.GetEmails).Methods(http.MethodGet)
	subrouter.HandleFunc("/customers/{id:[0-9,]+}/phones", controllers.GetPhones).Methods(http.MethodGet)

	return router
}
