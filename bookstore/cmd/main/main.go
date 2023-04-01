package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/saeedmdd/go-http-examples/pkg/routes"
	"log"
	"net/http"
)

const host = "127.0.0.1"
const port = "8080"

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	srv := host + ":" + port
	fmt.Printf("starting at http://%v\n", srv)
	log.Fatal(http.ListenAndServe(srv, r))
}
