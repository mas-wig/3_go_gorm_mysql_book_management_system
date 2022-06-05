package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mas-wig/3_go_mysql_book_management_system/src/model"
	"github.com/mas-wig/3_go_mysql_book_management_system/src/util"
)

var NewBook model.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	newBook := model.GetAllBook()
	res, _ := json.Marshal(newBook)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	arg := mux.Vars(r)
	bookId := arg["bookId"] // get argument id from /book/{bookId}
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}

	bookDetail, _ := model.GetBookById(ID)
	res, _ := json.Marshal(bookDetail)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &model.Book{}
	util.ParseBody(r, CreateBook)
	b := CreateBook.CreateBook()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DelateBook(w http.ResponseWriter, r *http.Request) {
	arg := mux.Vars(r)
	bookId := arg["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	model.DelateBook(ID)
	newBook := model.GetAllBook()
	res, _ := json.Marshal(newBook)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	updateBook := &model.Book{}
	util.ParseBody(r, updateBook)
	arg := mux.Vars(r)
	bookId := arg["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookDetail, db := model.GetBookById(ID)

	if updateBook.Author != "" &&
		updateBook.Country != "" &&
		updateBook.ImageLink != "" &&
		updateBook.Language != "" &&
		updateBook.Link != "" &&
		updateBook.Pages != 0 &&
		updateBook.Title != "" &&
		updateBook.Year != 0 {
		bookDetail.Author = updateBook.Author
		bookDetail.Country = updateBook.Country
		bookDetail.ImageLink = updateBook.ImageLink
		bookDetail.Language = updateBook.Language
		bookDetail.Link = updateBook.Link
		bookDetail.Pages = updateBook.Pages
		bookDetail.Title = updateBook.Title
		bookDetail.Year = updateBook.Year

	}

	db.Save(&bookDetail)
	res, _ := json.Marshal(bookDetail)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
