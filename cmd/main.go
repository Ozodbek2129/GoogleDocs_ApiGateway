package main

import (
	"api_gateway/api"
	"api_gateway/api/handler"
	"api_gateway/config"
	"log"
)

func main() {
	conf := config.Load()
	hand := handler.NewHandler()

	router := api.NewRouter(hand)
	log.Printf("server is running...")
	log.Fatal(router.Run(conf.API_GATEWAY))
}
