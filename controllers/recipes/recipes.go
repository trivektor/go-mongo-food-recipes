package recipes

import (
  "net/http"
  "encoding/json"
  "go-mongo-food/db"
  "go.mongodb.org/mongo-driver/bson"
  "go.mongodb.org/mongo-driver/bson/primitive"
)

type Recipe struct {
  Name string
  Description string
}

func Index(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("Recipes!\n"))
}

func Create(w http.ResponseWriter, r *http.Request) {
  var recipe Recipe

  json.NewDecoder(r.Body).Decode(&recipe)

  collection := db.Client.Database("food_recipes").Collection("recipes")
  res, _ := collection.InsertOne(db.Context, bson.M{"name": recipe.Name, "description": recipe.Description})
  id := res.InsertedID.(primitive.ObjectID)

  // https://stackoverflow.com/a/49937678/477697
  w.Write([]byte(id.Hex()))
}
