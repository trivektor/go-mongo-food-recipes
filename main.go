package main

import (
  "net/http"
  "github.com/gorilla/mux"
  "./controllers/recipes"
)

func main() {
  r := mux.NewRouter()
  r.HandleFunc("/recipes", recipes.Index)
  r.HandleFunc("/recipes", recipes.Create).Methods("POST")
  http.ListenAndServe(":8080", r)
}
