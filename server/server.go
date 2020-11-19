package server

import (
	"flag"
	"log"
	"net/http"
)

/* Begins the web server*/
func BeginServer() {

	var ip = flag.String("database_ip", "127.0.0.1", "the ip to the database")
	var port = flag.String("database_port", "5555", "the port to the database")
	var user = flag.String("database_username", "turbotage", "the username to the database")
	var password = flag.String("database_password", "1234", "the password to the database")
	var dbname = flag.String("database_name", "weatherstation", "the database name")

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
