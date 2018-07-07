package main

import (
	"log"
	"net/http"

	"WeatherGo/fetcher"

	"github.com/googollee/go-socket.io"
)

func runWebserver() {
	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}

	server.On("connection", func(so socketio.Socket) {
		log.Println("on connection")
		so.Join("chat")
		so.On("weather:query_windspeed", func(msg string) {
			log.Println(msg)
		})

		so.On("weather:query_windspeedmax", func(msg string) {

		})

		so.On("disconnection", func() {
			log.Println("on disconnect")
		})
	})

	server.On("error", func(so socketio.Socket, err error) {
		log.Println("error:", err)
	})

	http.Handle("/socket.io/", server)
	http.Handle("/", http.FileServer(http.Dir("./asset")))
	log.Println("Serving at localhost:5000...")
	log.Fatal(http.ListenAndServe(":5000", nil))
}

func main() {

	go fetcher.BeginFetch()

	runWebserver()
}
