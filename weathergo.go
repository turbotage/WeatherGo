package weather

import (
	"github.com/turbotage/WeatherGo/fetcher"
)

func main() {

	go fetcher.BeginFetching("Weather!212", "/dev/ttyACM0", 9600)

	//go server.BeginServer()
}
