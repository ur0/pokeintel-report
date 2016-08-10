package router

import (
	"github.com/gorilla/mux"
	"github.com/ur0/pokeintel-report/handlers"
)

func GetRouter() (r *mux.Router) {
	r = mux.NewRouter()
	r.HandleFunc("/", handlers.RootHandler)
  r.HandleFunc("/encounters", handlers.EncountersHandler).Methods("POST")
  
	return
}
