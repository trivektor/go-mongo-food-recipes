package recipes

import (
  "net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("Recipes!\n"))
}

func Create(w http.ResponseWriter, r *http.Request) {
  
}
