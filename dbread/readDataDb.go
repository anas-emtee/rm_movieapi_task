package dbread

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//MovieObjects is the Structure of movie object we will be saving in our database
type MovieObjects struct {
	ID            int    `json:"id"`
	Title         string `json:"title"`
	Description   string `json:"overview"`
	OriginalTitle string `json:"original_title"`
	PosterPath    string `json:"poster_path"`
}

//GetAllMovies is a function that collects all records from Database
func GetAllMovies() []byte {
	clientOptions := options.Client().ApplyURI("mongodb://127.0.0.1:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	collection := client.Database("rm_moviedb").Collection("movies")

	findOptions := options.Find()

	var results []*MovieObjects
	cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		var elem MovieObjects
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.TODO())
	fmt.Printf("Found multiple documents (array of pointers): %+v\n", results)

	imgbody, jsonErr := json.Marshal(results)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return imgbody
}

//FindMovie is a function that collects movie record from Database based on a given id
func FindMovie(f int) []byte {
	clientOptions := options.Client().ApplyURI("mongodb://127.0.0.1:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	collection := client.Database("rm_moviedb").Collection("movies")

	//filter := bson.D{{"id", f}}
	filter := bson.D{primitive.E{Key: "id", Value: f}}

	var result MovieObjects

	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Found a single document: %+v\n", result.Title)

	imgbody, jsonErr := json.Marshal(result)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return imgbody
}
