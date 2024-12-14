package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

/*
The json formatting in the struct fields is called struct field tags in Go. These tags are used to specify metadata for a struct field.
In this case, the tags specify how the field should be serialized or deserialized when converting the struct to/from JSON format.
postman the data is sent using the keywords written in the json.
*/
type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)

	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)

	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	var movie Movie

	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(100000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movies)

}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)

	for index, item := range movies {

		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)

			var updatedMovie Movie
			_ = json.NewDecoder(r.Body).Decode(&updatedMovie)

			updatedMovie.ID = params["id"]
			movies = append(movies, updatedMovie)

			json.NewEncoder(w).Encode(movies)
			return
		}
	}
}

func main() {
	fmt.Println("Hello!")

	movies = append(movies, Movie{
		ID:    "1",
		Isbn:  "12345",
		Title: "Inception",
		Director: &Director{
			FirstName: "Christopher",
			LastName:  "Nolan",
		},
	})

	movies = append(movies, Movie{
		ID:    "2",
		Isbn:  "67890",
		Title: "The Matrix",
		Director: &Director{
			FirstName: "Lana",
			LastName:  "Wachowski",
		},
	})
	r := mux.NewRouter()
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("starting port at port 8000")

	err := http.ListenAndServe(":8000", r)
	log.Fatal(err)
}
