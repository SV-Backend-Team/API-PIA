package suppliers

import (
	config "northwindApi/config"
)

type SupplierModel struct{}

//GET
func (supplierModel SupplierModel) GetSuppliers() ([]Supplier, error) {
	db, err := config.DB_export.OpenDB()
	if err != nil {
		return nil, err
	} else {
		var suppliers []Supplier
		db.Find(&suppliers)
		return suppliers, nil
	}
}

//GET{ID}
func (supplierModel SupplierModel) GetSupplierByID(id int) (Supplier, error) {
	db, err := config.DB_export.OpenDB()
	if err != nil {
		return Supplier{}, err
	} else {
		var suppliers Supplier
		db.Where("SupplierID = ?", id).Find(&suppliers) //Es literal el nombre de la columna
		return suppliers, nil
	}
}

//POST
func (supplierModel SupplierModel) CreateSuppliers(supplier *Supplier) error {
	db, err := config.DB_export.OpenDB()
	if err != nil {
		return err
	} else {
		db.Create(&supplier)
		return nil
	}
}

//PUT
func (supplierModel SupplierModel) UpdateSupplier(supplier *Supplier) error {
	db, err := config.DB_export.OpenDB()
	if err != nil {
		return err
	} else {
		db.Save(&supplier)
		return nil
	}
}

//Delete
func (supplierModel SupplierModel) DeleteSupplier(supplier Supplier) error {
	db, err := config.DB_export.OpenDB()
	if err != nil {
		return err
	} else {
		db.Delete(supplier)
		return nil
	}
}
