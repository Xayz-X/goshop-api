package main

import (
	"log"

	"github.com/Xayz-X/goshop-api/routes"
	"github.com/Xayz-X/goshop-api/services"
)

func main() {

	// run the server on port -> 3030
	server := services.NewAPIService(":3030")
	router := routes.GetRoutes()
	err := server.Run(router)
	if err != nil {
		log.Fatal("server start failed", err)
	}
}
