package services

import (
	"log"
	"net/http"
)

type APIService struct {
	addr string
}

func (s *APIService) Run(router *http.ServeMux) error {
	server := http.Server{
		Addr:    s.addr,
		Handler: router,
	}
	log.Println("server has started", s.addr)
	return server.ListenAndServe()

}

func NewAPIService(addr string) *APIService {
	return &APIService{
		addr: addr,
	}
}
