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
	r_emp := r.Group("/employee")
	r_emp.GET("/getemployees", employees.GetEmployees)
	r_emp.GET("/getemployee/:id", employees.GetEmployeeByID)
	r_emp.POST("/createemployee", employees.CreateEmployee)
	r_emp.PUT("/updateemployee", employees.UpdateEmployee)
	r_emp.DELETE("/deleteemployee/:id", employees.DeleteEmployeeByID)
}

func customerRoutes(r *echo.Group) {
	r_cus := r.Group("/customer")
	r_cus.GET("/getcustomers", customers.GetCutomers)
	r_cus.GET("/getcustomer/:id", customers.GetCustomerByID)
	r_cus.POST("/createcustomer", customers.CreateCustomer)
	r_cus.PUT("/updatecustomer", customers.UpdateCustomer)
	r_cus.DELETE("/deletecustomer/:id", customers.DeleteCustomerByID)
}

func supplierRoutes(r *echo.Group) {
	r_sup := r.Group("/supplier")
	r_sup.GET("/getsuppliers", suppliers.GetSuppliers)
	r_sup.GET("/getsupplier/:id", suppliers.GetSuppliersByID)
	r_sup.POST("/createsupplier", suppliers.CreateSuppliers)
	r_sup.PUT("/updatesupplier", suppliers.UpdateSupplier)
	r_sup.DELETE("/deletesupplier/:id", suppliers.DeleteSupplierByID)
}
