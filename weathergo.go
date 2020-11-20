package main

import (
	"flag"
	"fmt"
	"sync"

	"github.com/turbotage/WeatherGo/fetcher"
	"github.com/turbotage/WeatherGo/server"
)

func main() {

	var password = flag.String("database_password", "1234", "the password to the database")
	var serial_port = flag.String("serial_port", "/dev/ttyACM0", "the serial port to use for fetching")

	flag.Parse()

	var wg sync.WaitGroup

	fmt.Println("starting to fetch...")

	//doneFetching := make(chan bool, 1)
	wg.Add(1)
	go fetcher.BeginFetching(&wg, *password, *serial_port, 9600)

	wg.Add(1)
	go server.BeginServer(&wg, *password)

	fmt.Println("Main: Waiting for workers to finish")
	wg.Wait()
	fmt.Println("Main: Completed")

	//<-doneFetching
}
