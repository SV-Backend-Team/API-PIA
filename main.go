package main

import (
	"fmt"
	"net/http"
	config "northwindApi/config"
	routing "northwindApi/routing"

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

	routing.Router(router)

	//CORS
	err := http.ListenAndServe(":5000", handlers.CORS(handlers.AllowedHeaders([]string{"X-Request-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router))
	if err != nil {
		fmt.Println(err.Error())
	}
}
