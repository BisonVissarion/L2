package main

import (
	"dev11/handler"
	"dev11/middleware"
	"dev11/service"
	"log"
	"net/http"
)

func configureRoutes(serveMux *http.ServeMux, Server *handler.Server) {
	serveMux.HandleFunc("/create_event", Server.HandlerCreateEvent)
	serveMux.HandleFunc("/update_event", Server.HandlerUpdateEvent)
	serveMux.HandleFunc("/delete_event", Server.HandlerDeleteEvent)
	serveMux.HandleFunc("/events_for_day", Server.HandlerEventsForDay)
	serveMux.HandleFunc("/events_for_week", Server.HandlerEventsForWeek)
	serveMux.HandleFunc("/events_for_month", Server.HandlerEventsForMonth)
}

func main() {

	serveMux := http.NewServeMux()
	store := service.NewStore()
	storeServer := handler.NewServer(store)
	configureRoutes(serveMux, storeServer)

	login := middleware.Logging(serveMux)

	log.Fatal(http.ListenAndServe("localhost:8080", login))
}
