package models

import (
	"github.com/jinzhu/gorm"
	"github.com/trekhub/go-bookstore/pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string  `json:"name"`
	Author      string  `json:"author"`
	Publication string  `json:"publication"`
	Price       string `json:"price"`
}

// initialize the database
func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})

}

// create a book
func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

// get all books
func GetBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

//get a book
func GetBook(id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("id = ?", id).Find(&getBook)
	return &getBook, db
}

//delete a book
func DeleteBook(id int64) *Book {
	var book Book
	db.Where("id = ?", id).Delete(&book)
	return &book
}	