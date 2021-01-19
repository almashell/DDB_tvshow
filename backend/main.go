package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"time"
)

var client *mongo.Client

type TvShow struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name 	  string             `json:"name,omitempty" bson:"name,omitempty"`
}

type TvShowName struct {
	Name 	  string             `json:"name,omitempty" bson:"name,omitempty"`
}

func CreateTvshowEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var tvshow TvShowName
	_ = json.NewDecoder(request.Body).Decode(&tvshow)
	collection := client.Database("tvshow").Collection("tvshownames")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, _ := collection.InsertOne(ctx, tvshow)
	json.NewEncoder(response).Encode(result)
}

func GetTvshowEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	name, _ := params["name"]

	fmt.Println(name)

	var tvshow TvShow
	collection := client.Database("tvshow").Collection("tvshownames")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	err := collection.FindOne(ctx, TvShowName{Name: name}).Decode(&tvshow)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(tvshow)
}

func main() {
	fmt.Println("Starting the application...")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, _ = mongo.Connect(ctx, clientOptions)
	router := mux.NewRouter()
	router.HandleFunc("/api/save", CreateTvshowEndpoint).Methods("POST")
	router.HandleFunc("/api/find/{name}", GetTvshowEndpoint).Methods("GET")
	http.ListenAndServe(":12345", router)
}