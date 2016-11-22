package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/kapsteur/event.club/config"
	"github.com/kapsteur/event.club/controller"
	"log"
	"net/http"
)

func main() {
	conf := config.Conf()
	log.Printf("Event.club is started in %s:%d", conf.Env, conf.Port)

	r := mux.NewRouter()

	r.HandleFunc("/static/{file}", controller.StaticHandler)
	r.HandleFunc("/booking", controller.BookingHandler)
	r.HandleFunc("/", controller.HomeHandler)

	http.Handle("/", r)

	http.ListenAndServe(fmt.Sprintf(":%d", conf.Port), nil)
}
