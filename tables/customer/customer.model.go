package customers

import (
	config "northwindApi/config"
)

type CustomerModel struct {
}

//GET
func (customerModel CustomerModel) GetCutomers() ([]Customer, error) {
	db, err := config.DB_export.OpenDB()
	if err != nil {
		return nil, err
	} else {
		var customers []Customer
		db.Find(&customers)
		return customers, nil
	}
}

//GET{ID}
func (customerModel CustomerModel) GetCustomerByID(id string) (Customer, error) {
	db, err := config.DB_export.OpenDB()
	if err != nil {
		return Customer{}, err
	} else {
		var customer Customer
		db.Where("CustomerID like ?", id).Find(&customer) //Es literal el nombre de la columna
		return customer, nil
	}
}

//POST
func (customerModel CustomerModel) CreateCustomer(customer *Customer) error {
	db, err := config.DB_export.OpenDB()
	if err != nil {
		return err
	} else {
		db.Create(&customer)
		return nil
	}
}

//Update
func (customerModel CustomerModel) UpdateCustomer(customer *Customer) error {
	db, err := config.DB_export.OpenDB()
	if err != nil {
		return err
	} else {
		db.Save(&customer)
		return nil
	}
}

//Delete
func (customerModel CustomerModel) DeleteCustomer(customer Customer) error {
	db, err := config.DB_export.OpenDB()
	if err != nil {
		return err
	} else {
		db.Delete(customer)
		return nil
	}
}
