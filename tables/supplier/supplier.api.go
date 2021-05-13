package suppliers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

//GET
func GetSuppliers(c echo.Context) error {
	log.Println("select all")
	var supplierModel SupplierModel
	suppliers, err := supplierModel.GetSuppliers()
	if err != nil {
		log.Panicln(err)
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, suppliers)
}

//GET{ID}
func GetSuppliersByID(c echo.Context) error {
	log.Println("select by id")
	id := c.Param("id")
	supplierid, err := strconv.Atoi(id)
	if err != nil {
		log.Panicln(err)
		return c.String(http.StatusBadRequest, err.Error())
	}

	var supplierModel SupplierModel
	suppliers, err := supplierModel.GetSupplierByID(supplierid)
	if err != nil {
		log.Panicln(err)
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, suppliers)
}

//POST
func CreateSuppliers(c echo.Context) error {
	log.Println("create")
	var supplier Supplier
	err := json.NewDecoder(c.Request().Body).Decode(&supplier)
	if err != nil {
		log.Panicln(err)
		return c.String(http.StatusBadRequest, err.Error())
	}

	var supplierModel SupplierModel
	err2 := supplierModel.CreateSuppliers(&supplier)
	if err2 != nil {
		log.Panicln(err)
		return c.String(http.StatusInternalServerError, err2.Error())
	}
	return c.JSON(http.StatusOK, supplier)
}

//Update
func UpdateSupplier(c echo.Context) error {
	log.Println("update")
	var supplier Supplier
	err := json.NewDecoder(c.Request().Body).Decode(&supplier)
	if err != nil {
		log.Panicln(err)
		return c.String(http.StatusBadRequest, err.Error())
	}

	var supplierModel SupplierModel
	err2 := supplierModel.UpdateSupplier(&supplier)
	if err2 != nil {
		log.Panicln(err)
		return c.String(http.StatusInternalServerError, err2.Error())
	}
	return c.JSON(http.StatusOK, supplier)
}

//Delete
func DeleteSupplierByID(c echo.Context) error {
	log.Println("delete")
	id := c.Param("id")
	supplierid, _ := strconv.Atoi(id)
	var supplierModel SupplierModel
	suppliers, err := supplierModel.GetSupplierByID(supplierid)
	if err != nil {
		log.Panicln(err)
		return c.String(http.StatusBadRequest, err.Error())
	}

	err2 := supplierModel.DeleteSupplier(suppliers)
	if err2 != nil {
		log.Panicln(err)
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, suppliers)
}
