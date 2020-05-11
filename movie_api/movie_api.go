package main

import (
	"fmt"
	"log"

	"github.com/dgreat91/rm_movieapi_task/configuration"
)

func main() {
	cfg, err := configuration.NewConfig()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cfg.API.Key)
}
