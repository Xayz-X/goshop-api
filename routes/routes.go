package routes

import (
	"net/http"

	"github.com/Xayz-X/goshop-api/controllers"
)

func GetRoutes() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("/", healthCheck)
	router.HandleFunc("POST /user/register", controllers.UserRegisterHandler)
	return router
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Server is running..."))
}
