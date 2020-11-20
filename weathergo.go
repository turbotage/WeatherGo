package main

import (
	"fmt"
	"sync"

	"github.com/turbotage/WeatherGo/fetcher"
)

func main() {

	var wg sync.WaitGroup

	fmt.Println("starting to fetch...")

	doneFetching := make(chan bool, 1)
	wg.Add(1)
	go fetcher.BeginFetching(doneFetching, &wg, "Weather!212", "/dev/ttyACM0", 9600)

	fmt.Println("Main: Waiting for workers to finish")
	wg.Wait()
	fmt.Println("Main: Completed")

	<-doneFetching
}
