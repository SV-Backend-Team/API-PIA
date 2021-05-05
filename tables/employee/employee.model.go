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
	} else {
		var employees []Employee
		db.Find(&employees)
		return employees, nil
	}
}

//GET{ID}
func (employeeModel EmployeeModel) GetEmployeeByID(id int) (Employee, error) {
	db, err := config.DB_export.OpenDB()
	if err != nil {
		return Employee{}, err
	} else {
		var employees Employee
		db.Where("EmployeeID = ?", id).Find(&employees) //Es literal el nombre de la columna
		return employees, nil
	}
}

//POST
func (employeeModel EmployeeModel) CreateEmployee(employee *Employee) error {
	db, err := config.DB_export.OpenDB()
	if err != nil {
		return err
	} else {
		db.Create(&employee)
		return nil
	}
}

//Update
func (employeeModel EmployeeModel) UpdateEmployee(employee *Employee) error {
	db, err := config.DB_export.OpenDB()
	if err != nil {
		return err
	} else {
		db.Save(&employee)
		return nil
	}
}

//Delete
func (employeeModel EmployeeModel) DeleteEmployee(employee Employee) error {
	db, err := config.DB_export.OpenDB()
	if err != nil {
		return err
	} else {
		db.Delete(employee)
		return nil
	}
}
