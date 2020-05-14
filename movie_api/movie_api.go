package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/dgreat91/rm_movieapi_task/configuration"
	"github.com/dgreat91/rm_movieapi_task/filesave"
)

//title, description, filename and its original link
type Base struct {
	Page         int `json:"page"`
	TotalResults int `json:"total_results"`
	TotalPages   int `json:"total_pages"`
	Results      []struct {
		ID    int    `json:"id"`
		Title string `json:"title"`
		//Description   string
		OriginalTitle string `json:"original_title"`
		PosterPath    string `json:"poster_path"`
	}
}

type movieObjects struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	//Description   string
	OriginalTitle string `json:"original_title"`
	PosterPath    string `json:"poster_path"`
}

//Handle request /ping by responding with pong
func ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

//Handle Request /fetch To Fetch Movies From MOVIEDB API
func fetch(w http.ResponseWriter, r *http.Request) {
	apiCfg, err := configuration.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	apiKey := apiCfg.API.Key
	/*clientOptions := options.Client().ApplyURI("mongodb://127.0.0.1:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")*/
	//https://api.themoviedb.org/3/discover/movie?api_key="+api_key+"&language=en-US&sort_by=popularity.desc&include_adult=false&include_video=false&page=1
	url := "https://api.themoviedb.org/3/discover/movie?api_key=" + apiKey + "&language=en-US&sort_by=popularity.desc&include_adult=false&include_video=false&page=1"

	spaceClient := http.Client{
		Timeout: time.Second * 2, // Maximum of 2 secs
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "revenue-moster-task")

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	b := &Base{}
	jsonErr := json.Unmarshal(body, &b)

	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	fmt.Printf("%+v\n", b)
	fmt.Println(b.Results[0].ID)

	//var m

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers",
		"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	w.Write([]byte(body))

	movieJSON, jsonErr := json.Marshal(b.Results)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	sb := string(movieJSON)
	filesave.SaveData(sb)
}
func main() {
	cfg, err := configuration.NewConfig()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Server is running on '%s' port '%s'", cfg.Server.Host, cfg.Server.Port)

	//Request Handler For Request /ping
	http.HandleFunc("/ping", ping)
	http.HandleFunc("/fetch", fetch)

	//Start Web Server
	if err := http.ListenAndServe(cfg.Server.Host+":"+cfg.Server.Port, nil); err != nil {
		panic(err)
	}
}
