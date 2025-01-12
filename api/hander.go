package handler

import (
	"net/http"

	"github.com/Xayz-X/goshop-api/controllers"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetRoutes(dataBase *mongo.Database) *http.ServeMux {
	router := http.NewServeMux()

	userCol := controllers.NewUserCollection(dataBase.Collection("user"))

	// health check handler
	router.HandleFunc("GET /", Handler)

	// user handlers
	router.HandleFunc("POST /user/register", userCol.UserRegisterHandler)
	router.HandleFunc("DELETE /user", userCol.DeleteUserHandler)
	router.HandleFunc("GET /users", userCol.GetAllUserHandler)

	return router
}

func Handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Server is running..."))
}
