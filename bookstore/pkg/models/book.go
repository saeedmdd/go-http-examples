package models

import (
	"github.com/jinzhu/gorm"
	"github.com/saeedmdd/go-http-examples/pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:""json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}
func (b *Book) Create() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAll() []Book {
	var Books []Book
	db.Find(&Books)

	return Books
}

func FindById(Id uint64) (*Book, *gorm.DB) {
	var book Book
	db := db.Where("ID=?", Id).Find(&book)
	return &book, db
}

func Delete(Id uint64) Book {
	var book Book
	db.Where("ID=?", Id).Delete(book)
	return book
}

func (b *Book) Update(id uint64, book Book) (Book, *gorm.DB) {
	db = db.Model(b).Where("ID=?", id).Updates(book)
	return book, db
}
