package main

import (
	config "northwindApi/config"
	routing "northwindApi/routing"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	config.DB_export = config.DB{
		Server:   "localhost",
		Port:     1433,
		User:     "",
		Password: "",
		Database: "Northwind",
	}

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
	}))

	routing.Router(e)

	//CORS
	e.Logger.Fatal(e.Start(":5000"))

}
