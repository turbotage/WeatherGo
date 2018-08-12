package fetcher

import (
	"log"
	"time"

	"github.com/tarm/serial"
)

/* the type used to begin fetching */
type FetchingInfo struct {
	ip               string
	port             string
	user             string
	password         string
	dbname           string
	serialname       string
	baud             int
	rainupdatetime   int //in seconds
	windupdatetime   int //in seconds
	bme280updatetime int //in seconds
}

func serialReadLine() {

}

func rainFetch() {

}

func windUpdate() {

}

func gustUpdate() {

}

func bme280Update() {

}

func fetchCycle(fI FetchingInfo) {
	done := false
	for i := 1; done; i++ {
		if (i % fI.rainupdatetime) == 0 {
			rainFetch()
		}
		if (i % fI.windupdatetime) == 0 {
			windUpdate()
			gustUpdate()
		}
		if (i % fI.bme280updatetime) == 0 {
			bme280Update()
		}
		time.Sleep(time.Second)
	}
}

/* "BeginFetching the function used to begin fetching" */
func BeginFetching(fetchingInfo FetchingInfo) {

	c := &serial.Config{Name: fetchingInfo.serialname, Baud: fetchingInfo.baud}
	serialport, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(5 * time.Millisecond)

}
