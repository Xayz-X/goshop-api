// package main

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"net/http"

// 	"github.com/Xayz-X/goshop-api/controllers"
// 	"github.com/Xayz-X/goshop-api/routes"
// 	"github.com/Xayz-X/goshop-api/services"
// )

// func Hello(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello Kitty")
// }

// func main() {

// 	// connect to database
// 	client, database := services.ConnectToDb()
// 	defer client.Disconnect(context.Background())

// 	// create
// 	userCol := database.Collection("user")
// 	controllers.NewUserCollection(userCol)

// 	// run the server on port -> 3030
// 	server := services.NewAPIService(":3030")
// 	router := routes.GetRoutes(database)
// 	err := server.Run(router)
// 	if err != nil {
// 		log.Fatal("server start failed", err)
// 	}
// }

package main

import (
	"context"
	"log"
	"net/http"

	"github.com/Xayz-X/goshop-api/controllers"
	"github.com/Xayz-X/goshop-api/routes"
	"github.com/Xayz-X/goshop-api/services"
)

// Your entry point handler function
func Handler(w http.ResponseWriter, r *http.Request) {
	// connect to database
	client, database := services.ConnectToDb()
	defer client.Disconnect(context.Background())

	// create
	userCol := database.Collection("user")
	controllers.NewUserCollection(userCol)

	// create API routes and start the server
	server := services.NewAPIService(":3030")
	router := routes.GetRoutes(database)

	// run the server or return the handler to Vercel
	err := server.Run(router)
	if err != nil {
		log.Fatal("server start failed", err)
	}
}

// Vercel expects a handler function, so we return it here
func main() {
	http.HandleFunc("/", Handler)
}
