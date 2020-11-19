package main

import (
	"WeatherGo/fetcher"
)

func main() {

	go fetcher.BeginFetch()

	//go server.BeginServer()
}
