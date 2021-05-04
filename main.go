package main

import (
	"fmt"
	"net/http"
	config "northwindApi/config"
	employees "northwindApi/tables/employee"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	config.DB_export = config.DB{
		Server:   "localhost",
		Port:     1433,
		User:     "",
		Password: "",
		Database: "Northwind",
	}

	router := mux.NewRouter()
	router.HandleFunc("/api/employee/findall", employees.FindAll).Methods("GET")
	router.HandleFunc("/api/employee/search/{id}", employees.FindByID).Methods("GET")
	router.HandleFunc("/api/employee/create", employees.CreateEmployee).Methods("POST")
	router.HandleFunc("/api/employee/update", employees.UpdateEmployee).Methods("PUT")
	router.HandleFunc("/api/employee/delete/{id}", employees.DeleteByID).Methods("DELETE")
	err := http.ListenAndServe(":5000", handlers.CORS(handlers.AllowedHeaders([]string{"X-Request-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router))
	if err != nil {
		fmt.Println(err.Error())
	}
}
