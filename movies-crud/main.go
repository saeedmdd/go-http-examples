package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

const host = "localhost"
const port = "8080"

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func getMovies(writer http.ResponseWriter, request *http.Request) {
	headerType(writer)
	json.NewEncoder(writer).Encode(movies)
}

func headerType(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
}
func getMovie(writer http.ResponseWriter, request *http.Request) {
	headerType(writer)
	params := mux.Vars(request)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(writer).Encode(item)
			return
		}
	}
	http.Error(writer, "404 not found", http.StatusNotFound)
}

func createMovie(writer http.ResponseWriter, request *http.Request) {
	headerType(writer)
	var movie Movie
	json.NewDecoder(request.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(1000000))
	movies = append(movies, movie)
	json.NewEncoder(writer).Encode(movie)
}

func updateMovie(writer http.ResponseWriter, request *http.Request) {
	headerType(writer)
	params := mux.Vars(request)

	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			json.NewDecoder(request.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(writer).Encode(movie)
		}
	}
}

func deleteMovie(writer http.ResponseWriter, request *http.Request) {
	headerType(writer)
	params := mux.Vars(request)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	http.Error(writer, "404 not found", http.StatusNotFound)
}
func main() {
	r := mux.NewRouter()
	movies = append(movies, Movie{
		ID:    "1",
		Isbn:  "123123123",
		Title: "akbarian",
		Director: &Director{
			Firstname: "Akbar",
			Lastname:  "Abdi",
		},
	})
	movies = append(movies, Movie{
		ID:    "2",
		Isbn:  "456456456",
		Title: "asgharian",
		Director: &Director{
			Firstname: "Asghar",
			Lastname:  "Hemmat",
		},
	})
	movies = append(movies, Movie{
		ID:    "3",
		Isbn:  "789789789",
		Title: "mammadian",
		Director: &Director{
			Firstname: "Mammad",
			Lastname:  "Nobari",
		},
	})
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")
	serv := host + ":" + port
	fmt.Printf("starting at http://%v\n", serv)
	log.Fatal(http.ListenAndServe(serv, r))
}
