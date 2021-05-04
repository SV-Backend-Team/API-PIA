package employees

type Employee struct {
	EmployeeID      int    `gorm:"primary_Key;column:EmployeeID"` // Usar con precaucion
	LastName        string `gorm:"column:LastName"`
	FirstName       string `gorm:"column:FirstName"`
	Title           string `gorm:"column:Title"`
	TitleOfCourtesy string `gorm:"column:TitleOfCourtesy"`
	BirthDate       string `gorm:"column:BirthDate"`
	HireDate        string `gorm:"column:HireDate"`
	Address         string `gorm:"column:Address"`
	City            string `gorm:"column:City"`
	Region          string `gorm:"column:Region"`
	PostalCode      string `gorm:"column:PostalCode"`
	Country         string `gorm:"column:Country"`
	HomePhone       string `gorm:"column:HomePhone"`
	Extension       string `gorm:"column:Extension"`
	Notes           string `gorm:"column:Notes"`
	ReportsTo       string `gorm:"foreign_Key;column:ReportsTo"`
	PhotoPath       string `gorm:"column:PhotoPath"`
}

func (employee *Employee) TableName() string { return "Employees" }
