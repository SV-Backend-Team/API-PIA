package config

import (
	"fmt"

	"github.com/jinzhu/gorm" //ORM
	//_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type DB struct {
	Server   string
	Port     int
	User     string
	Password string
	Database string
}

var DB_export DB

func (db_connection DB) OpenDB() (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", "Northwind.db")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func (db_connection DB) ConnectToSQLServer() string {
	return fmt.Sprintf("server=%s; user id=%s; password=%s; port=%d; database=%s;",
		db_connection.User, db_connection.Password, db_connection.Server, db_connection.Port, db_connection.Database)
}

//db, err := gorm.Open("mssql", "sqlserver://username:password@localhost:1433? database=dbname")
