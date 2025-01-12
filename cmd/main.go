package main

import (
	"context"
	"log"

	handler "github.com/Xayz-X/goshop-api/api"
	"github.com/Xayz-X/goshop-api/controllers"
	"github.com/Xayz-X/goshop-api/services"
)

func main() {

	// connect to database
	client, database := services.ConnectToDb()
	defer client.Disconnect(context.Background())

	// create
	userCol := database.Collection("user")
	controllers.NewUserCollection(userCol)

	// run the server on port -> 3030
	server := services.NewAPIService(":3000")
	router := handler.GetRoutes(database)
	err := server.Run(router)
	if err != nil {
		log.Fatal("server start failed", err)
	}
}
