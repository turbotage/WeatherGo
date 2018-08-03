package fetcher

import (
	"flag"
	"time"
	"math"

	"github.com/cfreeman/embd"
	_ "github.com/cfreeman/embd/host/rpi" // This loads the RPi driver
)

const sleepTime = 5  //seconds
const sleepCycles = 60 //number of cycles

var anemometerPinNum uint8
var anemometerPin InterruptPin
var anemometerCount uint32

var rainbucketPinNum uint8
var rainbucketPin InterruptPin
var rainbucketCount uint32


//bme280 stuff
var bme280 BME80

var temperatureSum float32
var pressureSum float32
var humiditySum float32

//interrupts
func setupInterrupts() {
	//anemometer
	anemometerPin, err = embd.NewDigitalPin(anemometerPinNum)
	if err != nil {
		panic(err)
	}

	err = anemometerPin.SetDirection(embd.In)
	if err != nil {
		panic(err)
	}
	anemometerPin.ActiveLow(false)

	anemometerPin.Watch(embd.EdgeFalling, func() {
		anemometerCount++
	})

	//rainbucket
	rainbucketPin, err = embd.NewDigitalPin(rainbucketPinNum)
	if err != nil {
		panic(err)
	}

	err = rainbucketPin.SetDirection(embd.In)
	if err != nil {
		panic(err)
	}
	rainbucketPin.ActiveLow(false)

	rainbucketPin.Watch(embd.EdgeFalling, func() {
		rainbucketCount++
	}}
}

func endInterrupts() {
	anemometerPin.StopWatch()
	anemometerPin.Close()

	rainbucketPin.StopWatch()
	rainbucketPin.Close()
}




func zeroValues() {
	zeroWindDir()



	temperatureSum = 0
	pressureSum = 0
	humiditySum = 0

}

func cycleUpdate() {
	updateWindDirs()
}

func finalizeCycleResults() {
	degree, dirName := finalWindDir()
}

func BeginFetching() {
	flag.Parse()

	if err := embd.InitGPIO(); err != nil {
		panic(err)
	}
	defer embd.CloseGPIO()

	setupInterrupts()
	defer endInterrupts()

	for j := 0; true; j++ {
		//one cycle
		zeroValues()
		for i := 0; i < 60; i++ {
			cycleUpdate()
		}
		finalizeCycleResults()
	}


}
