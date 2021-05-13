package employees

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

//GET
func GetEmployees(c echo.Context) error {
	log.Println("select all")
	var employeeModel EmployeeModel
	employees, err := employeeModel.GetEmployees()
	if err != nil {
		log.Panicln(err)
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, employees)
}

//GET{ID}
func GetEmployeeByID(c echo.Context) error {
	log.Println("select by id")
	id := c.Param("id")
	employeeid, err := strconv.Atoi(id)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	var employeeModel EmployeeModel
	employees, err := employeeModel.GetEmployeeByID(employeeid)
	if err != nil {
		log.Panicln(err)
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, employees)
}

//POST
func CreateEmployee(c echo.Context) error {
	log.Println("create")
	var employee Employee
	err := json.NewDecoder(c.Request().Body).Decode(&employee)
	if err != nil {
		log.Panicln(err)
		return c.String(http.StatusBadRequest, err.Error())
	}

	var employeeModel EmployeeModel
	err = employeeModel.CreateEmployee(&employee)
	if err != nil {
		log.Panicln(err)
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, employee)
}

//Update
func UpdateEmployee(c echo.Context) error {
	log.Println("update")
	var employee Employee
	err := json.NewDecoder(c.Request().Body).Decode(&employee)
	if err != nil {
		log.Panicln(err)
		return c.String(http.StatusBadRequest, err.Error())
	}

	var employeeModel EmployeeModel
	err = employeeModel.UpdateEmployee(&employee)
	if err != nil {
		log.Panicln(err)
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, employee)
}

//Delete
func DeleteEmployeeByID(c echo.Context) error {
	log.Println("delete")
	id := c.Param("id")
	employeeid, _ := strconv.Atoi(id)
	var employeeModel EmployeeModel
	employees, err := employeeModel.GetEmployeeByID(employeeid)
	if err != nil {
		log.Panicln(err)
		return c.String(http.StatusBadRequest, err.Error())
	}

	err = employeeModel.DeleteEmployee(employees)
	if err != nil {
		log.Panicln(err)
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, employees)
}
