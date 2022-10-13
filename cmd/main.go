package main

import (
	"customerapp/mapstore"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func main() {
	fmt.Println("Welcome to Customer App RestAPI")
	logger, _ := zap.NewProduction()
	inmemrepo, err := mapstore.NewMapStore()
	if err != nil {
		log.Fatal("Error:", err)
	}
	h := &CustomerHandler{
		repo:   inmemrepo,
		Logger: logger,
	}
	router := initializeRoutes(h)
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	log.Println("Listening...")
	server.ListenAndServe()

}

func initializeRoutes(h *CustomerHandler) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/customer", h.GetAll).Methods("GET")
	r.HandleFunc("/api/customer/{custid}", h.Get).Methods("GET")
	r.HandleFunc("/api/customer", h.Post).Methods("POST")
	r.HandleFunc("/api/customer/{custid}", h.Put).Methods("PUT")
	r.HandleFunc("/api/customer/{custid}", h.Delete).Methods("DELETE")
	return r
}
