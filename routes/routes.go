package routes

import (
	"net/http"

	"github.com/Xayz-X/goshop-api/controllers"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetRoutes(dataBase *mongo.Database) *http.ServeMux {
	router := http.NewServeMux()

	userCol := controllers.NewUserCollection(dataBase.Collection("user"))

	router.HandleFunc("GET /", healthCheck)
	router.HandleFunc("POST /user/register", userCol.UserRegisterHandler)
	router.HandleFunc("GET /users", userCol.ListAllUser)

	return router
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Server is running..."))
}
