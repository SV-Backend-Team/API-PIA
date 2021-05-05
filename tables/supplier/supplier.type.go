package suppliers

type Supplier struct {
	SupplierID   int    `gorm:"primaryKey;column:SupplierID"`
	CompanyName  string `gorm:"column:CompanyName"`
	ContactName  string `gorm:"column:ContactName"`
	ContactTitle string `gorm:"column:ContactTitle"`
	Address      string `gorm:"column:Address"`
	City         string `gorm:"column:City"`
	Region       string `gorm:"column:Region"`
	PostalCode   string `gorm:"column:PostalCode"`
	Country      string `gorm:"column:Country"`
	Phone        string `gorm:"column:Phone"`
	Fax          string `gorm:"column:Fax"`
	HomePage     string `gorm:"column:HomePage"`
}

func (supplier *Supplier) TableName() string { return "Suppliers" }
