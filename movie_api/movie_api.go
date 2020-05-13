package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dgreat91/rm_movieapi_task/configuration"
)

//title, description, filename and its original link
type movieObjects struct {
	ID          int
	Title       string
	Description string
	Filename    string
	PosterURL   string
}

//Handle request /ping by responding with pong
func ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}
func main() {
	cfg, err := configuration.NewConfig()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Server is running on '%s' port '%s'", cfg.Server.Host, cfg.Server.Port)

	//Request Handler For Request /ping
	http.HandleFunc("/ping", ping)

	//Start Web Server
	if err := http.ListenAndServe(cfg.Server.Host+":"+cfg.Server.Port, nil); err != nil {
		panic(err)
	}
}
