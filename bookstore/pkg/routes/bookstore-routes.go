package routes

import (
	"github.com/gorilla/mux"
	"github.com/saeedmdd/go-http-examples/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBookById).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBookById).Methods("DELETE")
}
