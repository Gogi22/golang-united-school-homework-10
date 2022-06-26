package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

func Start(host string, port int) {
	router := mux.NewRouter()

	router.HandleFunc("/name/{param}", GetByParam).Methods(http.MethodGet)
	router.HandleFunc("/bad", Bad).Methods(http.MethodGet)
	router.HandleFunc("/data", Data).Methods(http.MethodPost)
	router.HandleFunc("/headers", Headers).Methods(http.MethodPost)

	log.Println(fmt.Printf("Starting API server on %s:%d\n", host, port))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router); err != nil {
		log.Fatal(err)
	}
}

func GetByParam(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.Write([]byte(fmt.Sprintf("Hello, %s!", vars["param"])))
}

func Bad(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(500)
}

func Data(w http.ResponseWriter, r *http.Request) {
	data, _ := io.ReadAll(r.Body)
	w.Write([]byte(fmt.Sprintf("I got message:\n%s", data)))
}

func Headers(w http.ResponseWriter, r *http.Request) {
	a, _ := strconv.Atoi(r.Header.Get("a"))
	b, _ := strconv.Atoi(r.Header.Get("b"))
	w.Header().Set("a+b", strconv.Itoa(a+b))
}

func main() {
	host := os.Getenv("HOST")
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8081
	}
	Start(host, port)
}
