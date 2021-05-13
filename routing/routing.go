package routing

import (
	customers "northwindApi/tables/customer"
	employees "northwindApi/tables/employee"
	suppliers "northwindApi/tables/supplier"

	"github.com/labstack/echo/v4"
)

func Router(e *echo.Echo) {
	r := e.Group("/api")
	employeeRoutes(r)
	customerRoutes(r)
	supplierRoutes(r)
}

func employeeRoutes(r *echo.Group) {
	r.GET("/employee/getemployees", employees.GetEmployees)
	r.GET("/employee/getemployee/:id", employees.GetEmployeeByID)
	r.POST("/employee/createemployee", employees.CreateEmployee)
	r.PUT("/employee/updateemployee", employees.UpdateEmployee)
	r.DELETE("/employee/deleteemployee/:id", employees.DeleteEmployeeByID)
}

func customerRoutes(r *echo.Group) {
	r.GET("/customer/getcustomers", customers.GetCutomers)
	r.GET("/customer/getcustomers/:id", customers.GetCustomerByID)
	r.POST("/customer/createcustomer", customers.CreateCustomer)
	r.PUT("/customer/updatecustomer", customers.UpdateCustomer)
	r.DELETE("/customer/deletecustomer/:id", customers.DeleteCustomerByID)
}

func supplierRoutes(r *echo.Group) {
	r.GET("/supplier/getsuppliers", suppliers.GetSuppliers)
	r.GET("/supplier/getsuppliers/:id", suppliers.GetSuppliersByID)
	r.POST("/supplier/createsupplier", suppliers.CreateSuppliers)
	r.PUT("/supplier/updatesupplier", suppliers.UpdateSupplier)
	r.DELETE("/supplier/deletesupplier/:id", suppliers.DeleteSupplierByID)
}
