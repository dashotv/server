package models

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/lib/pq"
)

var DB *gorm.DB
var Driver = "postgres"
var Datasource = "postgresql://dashotv@localhost:26257/dashotv_development?sslmode=disable"

func init() {
	var err error
	DB, err = gorm.Open(Driver, Datasource)
	if err != nil {
		log.Fatal(err)
	}

	DB.LogMode(false)
	DB.AutoMigrate(&Release{})
}

type Base struct{}

func IsFatalError(err error) bool {
	if perr, ok := err.(*pq.Error); ok && perr.Code == "23505" {
		return false
	}
	return true
}
