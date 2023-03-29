package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = "8080"
const host = "localhost"

func helloHandler(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/hello" {
		http.Error(writer, "404 not found", http.StatusNotFound)
		return
	}
	if request.Method != "GET" {
		http.Error(writer, "request not supported", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(writer, "hello!")
}

func formHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "POST" {
		http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	if err := request.ParseForm(); err != nil {
		fmt.Fprintf(writer, "error is %v", err)
		return
	}
	fmt.Fprintf(writer, "POST request is successful\n")
	name := request.FormValue("name")
	address := request.FormValue("address")
	fmt.Fprintf(writer, "name is %s\n", name)
	fmt.Fprintf(writer, "address is %s\n", address)
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/form", formHandler)
	serv := host + ":" + port
	fmt.Printf("starting at http://%v\n", serv)
	if err := http.ListenAndServe(serv, nil); err != nil {
		log.Fatal(err)
	}
}
