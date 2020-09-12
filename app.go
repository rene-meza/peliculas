package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	. "./config"
	. "./dao"
	"github.com/gorilla/mux"
)

var dao = MoviesDAO{}
var config = Config{}

//GET
//ENDPOINT: http:localhost:8080/movies
func allMoviesEndPoint(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintln(w, "Muestras todas las peliculas")
	movies, err := dao.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, movies)
}

//GET
//ENDPOINT: http:localhost:8080/movies/{id}
//ENDPOINT: http:localhost:8080/movies/4
func findMovieEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Muestra pelicula específica segun id")
}

// POST
// ENDPOINT: http:localhost:8080/movies
func createMovieEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Registra una película")
}

// PUT
// ENDPOINT: http:localhost:8080/movies
func updateMovieEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Actualiza una película")
}

// DELETE
// ENDPOINT: http:localhost:8080/movies/{id}
func deleteMovieEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Elimina una película segun id")
}

func init() {
	config.Read()
	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
}
func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}
func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
func main() {
	r := mux.NewRouter()

	r.HandleFunc("/movies", allMoviesEndPoint).Methods("GET")
	r.HandleFunc("/movies/{ID}", findMovieEndPoint).Methods("GET")
	r.HandleFunc("/movies", createMovieEndPoint).Methods("POST")
	r.HandleFunc("/movies", updateMovieEndPoint).Methods("UPDATE")
	r.HandleFunc("/movies/{ID}", deleteMovieEndPoint).Methods("DELETE")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
