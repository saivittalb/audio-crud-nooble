package router

import (
	"audio-crud-nooble/middleware"

	"github.com/gorilla/mux"
)

// Router is exported and used in main.go
func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/short/{id}", middleware.GetShort).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/shorts", middleware.GetAllShorts).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/newshort", middleware.CreateShort).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/shorts/{id}", middleware.UpdateShort).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/deleteshort/{id}", middleware.DeleteShort).Methods("DELETE", "OPTIONS")

	return router
}
