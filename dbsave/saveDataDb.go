package dbsave

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

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

//SaveDataDb is a function that saves our result to database
func SaveDataDb(s string) {
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

	var movs []MovieObjects
	jsonErr := json.Unmarshal([]byte(s), &movs)

	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	for k := range movs {
		fmt.Printf("The image '%s'\n", movs[k].Title)
		insertResult, err := collection.InsertOne(context.TODO(), movs[k])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Inserted a Single Document: ", insertResult.InsertedID)

		//fmt.Printf( "The image '%s' is located at '%s'\n", imgs[k].Title, imgs[k].Url );
		//add imgs[k] to mongodb
	}
}
