package customers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

//GET
func GetCutomers(c echo.Context) error {
	log.Println("select all")
	var customerModel CustomerModel
	customers, err := customerModel.GetCutomers()
	if err != nil {
		log.Panicln(err)
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, customers)
}

//GET{ID}
func GetCustomerByID(c echo.Context) error {
	log.Println("select by id")
	id := c.Param("id")
	var customerModel CustomerModel
	customers, err := customerModel.GetCustomerByID(id)
	if err != nil {
		log.Panicln(err)
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, customers)
}

//POST
func CreateCustomer(c echo.Context) error {
	log.Println("create")
	var customer Customer
	err := json.NewDecoder(c.Request().Body).Decode(&customer)
	if err != nil {
		log.Panicln(err)
		return c.String(http.StatusBadRequest, err.Error())
	}

	var customerModel CustomerModel
	err = customerModel.CreateCustomer(&customer)
	if err != nil {
		log.Panicln(err)
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, customer)
}

//Update
func UpdateCustomer(c echo.Context) error {
	log.Println("update")
	var customer Customer
	err := json.NewDecoder(c.Request().Body).Decode(&customer)
	if err != nil {
		log.Panicln(err)
		return c.String(http.StatusBadRequest, err.Error())
	}
	if customer.CustomerID == "" {
		err_msg := "Debe incluir la ID"
		log.Panicln(err_msg)
		return c.String(http.StatusBadRequest, err_msg)
	}
	var customerModel CustomerModel
	err = customerModel.UpdateCustomer(&customer)
	if err != nil {
		log.Panicln(err)
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, customer)
}

//Delete
func DeleteCustomerByID(c echo.Context) error {
	log.Println("delete")
	id := c.Param("id")
	var customerModel CustomerModel
	customers, err := customerModel.GetCustomerByID(id)
	if err != nil {
		log.Panicln(err)
		return c.String(http.StatusBadRequest, err.Error())
	}
	err = customerModel.DeleteCustomer(customers)
	if err != nil {
		log.Panicln(err)
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, customers)
}
