package employees

import (
	"database/sql"
)

type Employee struct {
	EmployeeID      int            `gorm:"primary_Key;column:EmployeeID"` // Usar con precaucion
	LastName        string         `gorm:"column:LastName"`
	FirstName       string         `gorm:"column:FirstName"`
	Title           sql.NullString `gorm:"column:Title"`
	TitleOfCourtesy sql.NullString `gorm:"column:TitleOfCourtesy"`
	BirthDate       sql.NullString `gorm:"column:BirthDate"`
	HireDate        sql.NullString `gorm:"column:HireDate"`
	Address         sql.NullString `gorm:"column:Address"`
	City            sql.NullString `gorm:"column:City"`
	Region          sql.NullString `gorm:"column:Region"`
	PostalCode      sql.NullString `gorm:"column:PostalCode"`
	Country         sql.NullString `gorm:"column:Country"`
	HomePhone       sql.NullString `gorm:"column:HomePhone"`
	Extension       sql.NullString `gorm:"column:Extension"`
	Notes           sql.NullString `gorm:"column:Notes"`
	ReportsTo       sql.NullString `gorm:"foreignKey;column:ReportsTo"`
	PhotoPath       sql.NullString `gorm:"column:PhotoPath"`
}

func (employee *Employee) TableName() string { return "Employees" }
