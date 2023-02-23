package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

// helps to connect to the database
func Connect() {
	d, err := gorm.Open("mysql", "[username]:[password]@/golang?charset=utf8&parseTime=True&loc=Local")

	//checking for errors
	if err != nil {
		fmt.Println(err)

		panic("failed to connect database")

	}
	db = d
}

// Return a pointer to our database connection
func GetDB() *gorm.DB {
	return db
}
