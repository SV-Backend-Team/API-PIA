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
	router.HandleFunc("/api/employee/findall", employees.FindAll).Methods("GET")
	router.HandleFunc("/api/employee/search/{id}", employees.FindByID).Methods("GET")
	router.HandleFunc("/api/employee/create", employees.CreateEmployee).Methods("POST")
	router.HandleFunc("/api/employee/update", employees.UpdateEmployee).Methods("PUT")
	router.HandleFunc("/api/employee/delete/{id}", employees.DeleteByID).Methods("DELETE")
}

func customerRoutes(router *mux.Router) {
	router.HandleFunc("/api/customer/getcustomers", customers.GetCutomers).Methods("GET")
	router.HandleFunc("/api/customer/getcustomers/{id}", customers.GetCustomerByID).Methods("GET")
	router.HandleFunc("/api/customer/createcustomer", customers.CreateCustomer).Methods("POST")
	router.HandleFunc("/api/customer/updatecustomer", customers.UpdateCustomer).Methods("PUT")
	router.HandleFunc("/api/customer/deletecustomer/{id}", customers.DeleteCustomerByID).Methods("DELETE")
}

func supplierRoutes(router *mux.Router) {
	router.HandleFunc("/api/customer/getsuppliers", suppliers.GetSuppliers).Methods("GET")
	router.HandleFunc("/api/customer/getsuppliers/{id}", suppliers.GetSuppliersByID).Methods("GET")
	router.HandleFunc("/api/customer/createsupplier", suppliers.CreateSuppliers).Methods("POST")
	router.HandleFunc("/api/customer/updatesupplier", suppliers.UpdateSupplier).Methods("PUT")
	router.HandleFunc("/api/customer/deletesupplier/{id}", suppliers.DeleteSupplierByID).Methods("DELETE")
}
