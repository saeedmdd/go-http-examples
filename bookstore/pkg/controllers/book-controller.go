package controllers

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/saeedmdd/go-http-examples/pkg/models"
	"github.com/saeedmdd/go-http-examples/pkg/utils"
	"net/http"
	"strconv"
)

var NewBook models.Book

func CreateBook(writer http.ResponseWriter, request *http.Request) {
	book := &models.Book{}
	utils.ParseBody(request, book)
	utils.ResponseBody(writer, book.Create(), "ok", "book created", http.StatusCreated)
}

func GetBooks(writer http.ResponseWriter, request *http.Request) {
	newBooks := models.GetAll()
	utils.ResponseBody(writer, newBooks, "ok", "all books", http.StatusOK)
}

func GetBookById(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	bookId := params["bookId"]
	ID, err := strconv.ParseUint(bookId, 0, 0)
	if err != nil {
		fmt.Println(err)
	}
	bookDetails, _ := models.FindById(ID)
	if bookDetails.ID == 0 {
		http.Error(writer, "404 not found", http.StatusNotFound)
		return
	}
	utils.ResponseBody(writer, bookDetails, "ok", "all books", http.StatusOK)
}
func UpdateBookById(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	bookId := params["bookId"]
	ID, err := strconv.ParseUint(bookId, 0, 0)
	if err != nil {
		fmt.Println(err)
	}
	updateBook := models.Book{}
	utils.ParseBody(request, &updateBook)
	updateBook.ID = uint(ID)
	newBook, _ := updateBook.Update(ID, updateBook)
	utils.ResponseBody(writer, newBook, "ok", "book updated", http.StatusOK)
}

func DeleteBookById(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("delete")
	params := mux.Vars(request)
	bookId := params["bookId"]
	ID, err := strconv.ParseUint(bookId, 0, 0)
	if err != nil {
		fmt.Println(err)
	}
	book := models.Delete(ID)
	utils.ResponseBody(writer, book, "ok", "book deleted", http.StatusOK)
}
