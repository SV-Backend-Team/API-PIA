package employees

import (
	config "northwindApi/config"
)

type EmployeeModel struct{}

//GET
func (employeeModel EmployeeModel) GetEmployees() ([]Employee, error) {
	db, err := config.DB_export.OpenDB()
	if err != nil {
		return nil, err
	}
	var employees []Employee
	db.Find(&employees)
	return employees, nil

}

//GET{ID}
func (employeeModel EmployeeModel) GetEmployeeByID(id string) (Employee, error) {
	db, err := config.DB_export.OpenDB()
	if err != nil {
		return Employee{}, err
	}
	var employees Employee
	db.Where("EmployeeID = ?", id).Find(&employees) //Es literal el nombre de la columna
	return employees, nil

}

//POST
func (employeeModel EmployeeModel) CreateEmployee(employee *Employee) error {
	db, err := config.DB_export.OpenDB()
	if err != nil {
		return err
	}
	// Buscar una forma de hacer select de todos los campos de forma automática(?)
	db.Select("FirstName", "LastName").Create(&employee)
	return nil

}

//PUT
func (employeeModel EmployeeModel) UpdateEmployee(employee *Employee) error {
	db, err := config.DB_export.OpenDB()
	if err != nil {
		return err
	}
	db.Select("LastName", "FirstName").Save(&employee)
	return nil

}

//DELETE
func (employeeModel EmployeeModel) DeleteEmployee(employee Employee) error {
	db, err := config.DB_export.OpenDB()
	if err != nil {
		return err
	}
	db.Delete(employee)
	return nil

}
