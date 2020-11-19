package main

import (
	"sync"

	"github.com/turbotage/WeatherGo/fetcher"
)

func main() {

	var wg sync.WaitGroup

	go fetcher.BeginFetching(&wg, "Weather!212", "/dev/ttyACM0", 9600)

	//go server.BeginServer()
}
