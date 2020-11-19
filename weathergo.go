package weather

import (
	"WeatherGo/fetcher"
)

func main() {

	go fetcher.BeginFetch("Weather!212", "/dev/ttyACM0", 9600)

	//go server.BeginServer()
}
