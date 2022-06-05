package routes

import (
	//"github.com/mas-wig/3_go_mysql_book_management_system/src/controller"
	"github.com/gorilla/mux"
	"github.com/mas-wig/3_go_mysql_book_management_system/src/controller"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controller.GetBook).Methods("GET")
	router.HandleFunc("/book/", controller.CreateBook).Methods("POST")
	router.HandleFunc("/book/{bookId}", controller.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", controller.DelateBook).Methods("DELETE")
	router.HandleFunc("/book/{bookId}", controller.UpdateBook).Methods("PUT")
}
