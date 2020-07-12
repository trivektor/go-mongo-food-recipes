package main

import (
  "context"
  "time"
  "net/http"
  "os"
  "github.com/gorilla/mux"
  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
  "go-mongo-food/controllers/recipes"
  "go-mongo-food/db"
)

func main() {
  db.Client, _ = mongo.NewClient(options.Client().ApplyURI(os.Getenv("CHATTYY_DB_CONNECTION_STRING")))
  db.Context, db.CancelFunc = context.WithTimeout(context.Background(), 10*time.Second)
  db.Client.Connect(db.Context)

  r := mux.NewRouter()
  r.HandleFunc("/recipes", recipes.Create).Methods("POST")
  r.HandleFunc("/recipes", recipes.Index)
  http.ListenAndServe(":8080", r)

  defer db.Client.Disconnect(db.Context)
  defer db.CancelFunc()
}
