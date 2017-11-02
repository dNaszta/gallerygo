package main

import (
	"fmt"
	"gopkg.in/gorilla/mux.v1"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprint(w, "{\"page\" : \"Home\"}")
}

func init() {
	Load()
	fmt.Println("Configs:", Configs.toString())
	RunMongo()
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	http.ListenAndServe(Configs.Port, r)
}