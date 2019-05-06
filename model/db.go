package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)
var(
	DB *gorm.DB
)
func init() {
	db, err := gorm.Open("postgres",  "dbname=test")
	if err != nil{
		panic(err)
	}


	DB = db
}
