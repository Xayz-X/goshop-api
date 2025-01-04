package routes

import "net/http"

func GetRoutes() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("/", healthCheck)
	return router
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Server is running..."))
}
