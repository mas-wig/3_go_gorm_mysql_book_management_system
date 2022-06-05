package model

import (
	"github.com/jinzhu/gorm"
	"github.com/mas-wig/3_go_mysql_book_management_system/src/config"
)

var db *gorm.DB = config.Connection()

type Book struct {
	gorm.Model
	Author    string `gorm:"" json:"author"`
	Country   string `json:"country"`
	ImageLink string `json:"imageLink"`
	Language  string `json:"language"`
	Link      string `json:"link"`
	Title     string `json:"title"`
	Year      int64  `gorm:"type:int(10)" json:"year"`
	Pages     int64  `gorm:"type:int(10)"`
}

func init() {
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBook() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

func DelateBook(ID int64) Book {
	var book Book

	db.Where("ID=?", ID).Delete(book)
	return book
}
