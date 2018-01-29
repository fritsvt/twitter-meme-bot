package web

import (
	"github.com/gorilla/mux"
	"twitter-meme-bot/web/controllers"
)

func RegisterRoutes(r mux.Router) (*mux.Router) {
	r.HandleFunc("/", controllers.GetHome).Methods("GET")

	r.HandleFunc("/schedule", controllers.GetNewSchedule).Methods("GET")
	r.HandleFunc("/schedule", controllers.PostNewSchedule).Methods("POST")
	r.HandleFunc("/schedule/delete", controllers.GetDeleteSchedule).Methods("GET")

	return &r
}
