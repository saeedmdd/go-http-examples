package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	_ "net/http/pprof"
)

type Book struct {
	Isbn string `json:"isbn"`
	Name string `json:"name"`
}

var BookList []Book = []Book{}

func main() {
	mx := http.DefaultServeMux
	mx.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		x := make(map[string]string)
		x["mammad"] = "mammad"
		x["asghar"] = "asghar"
		jsonVar, _ := json.Marshal(x)
		writer.Header().Add("content-type", "application/json")
		writer.WriteHeader(200)
		writer.Write(jsonVar)
	})
	mx.Handle("/book", http.HandlerFunc(handleBook))
	http.ListenAndServe("localhost:8080", nil)
}

func handleBook(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		fmt.Println("get req")
		books, err := json.Marshal(BookList)
		if err == nil {
			writer.Write(books)
			return
		}
		panic(err)
		return
	} else if request.Method == "POST" {
		newBook := Book{}
		body, err := io.ReadAll(request.Body)

		if err == nil {
			json.Unmarshal(body, &newBook)
			fmt.Println(newBook)

			for _, book := range BookList {
				if book.Isbn == newBook.Isbn {
					writer.WriteHeader(http.StatusBadRequest)
					return
				}
			}
			BookList = append(BookList, newBook)
			writer.WriteHeader(http.StatusCreated)
			return
		}
		panic(err)
	}
	writer.WriteHeader(http.StatusBadRequest)
	return

}
