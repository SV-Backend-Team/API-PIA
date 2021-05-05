package routing

import (
	customers "northwindApi/tables/customer"
	employees "northwindApi/tables/employee"
	suppliers "northwindApi/tables/supplier"

	"github.com/gorilla/mux"
)

func Router(router *mux.Router) {
	employeeRoutes(router)
	customerRoutes(router)
	supplierRoutes(router)
}

func employeeRoutes(router *mux.Router) {
	router.HandleFunc("/api/employee/getemployees", employees.GetEmployees).Methods("GET")
	router.HandleFunc("/api/employee/getemployee/{id}", employees.GetEmployeeByID).Methods("GET")
	router.HandleFunc("/api/employee/createemployee", employees.CreateEmployee).Methods("POST")
	router.HandleFunc("/api/employee/updateemployee", employees.UpdateEmployee).Methods("PUT")
	router.HandleFunc("/api/employee/deleteemployee/{id}", employees.DeleteEmployeeByID).Methods("DELETE")
}

func customerRoutes(router *mux.Router) {
	router.HandleFunc("/api/customer/getcustomers", customers.GetCutomers).Methods("GET")
	router.HandleFunc("/api/customer/getcustomers/{id}", customers.GetCustomerByID).Methods("GET")
	router.HandleFunc("/api/customer/createcustomer", customers.CreateCustomer).Methods("POST")
	router.HandleFunc("/api/customer/updatecustomer", customers.UpdateCustomer).Methods("PUT")
	router.HandleFunc("/api/customer/deletecustomer/{id}", customers.DeleteCustomerByID).Methods("DELETE")
}

func supplierRoutes(router *mux.Router) {
	router.HandleFunc("/api/supplier/getsuppliers", suppliers.GetSuppliers).Methods("GET")
	router.HandleFunc("/api/supplier/getsuppliers/{id}", suppliers.GetSuppliersByID).Methods("GET")
	router.HandleFunc("/api/supplier/createsupplier", suppliers.CreateSuppliers).Methods("POST")
	router.HandleFunc("/api/supplier/updatesupplier", suppliers.UpdateSupplier).Methods("PUT")
	router.HandleFunc("/api/supplier/deletesupplier/{id}", suppliers.DeleteSupplierByID).Methods("DELETE")
}
