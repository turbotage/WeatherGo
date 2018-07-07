package fetcher

import (
	"flag"
	"time"

	"github.com/cfreeman/embd"
	_ "github.com/cfreeman/embd/host/rpi" // This loads the RPi driver
)

const sleepTime = 5  //seconds
const sleepCycles = 60 //number of cycles

type vec2 struct {
	x float32
	y float32
}

var anemometerPinNum uint8
var anemometerPin InterruptPin
var anemometerCount uint32

var rainbucketPinNum uint8
var rainbucketPin InterruptPin
var rainbucketCount uint32

var winddirPin AnalogPin
var winddirSum vec2

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
		anemometerCount++
	}}
}

func endInterrupts() {
	anemometerPin.StopWatch()
	anemometerPin.Close()

	rainbucketPin.StopWatch()
	rainbucketPin.Close()
}

//winddirs
func updateWindDirs() {
	//"N", "NNO", "NO", "ONO", "O", "OSO", "SO", "SSO", "S", "SSW", "SW", "WSW", "W", "WNW", "NW", "NNW"
	dirValues := [16]int{ 788, 407, 464, 83, 92, 65, 185, 126, 288, 245, 634, 603, 947, 830, 889, 706 }
	dirDegrees := [16]float32{
		22.5 * 0, 22.5 * 1, 22.5 * 2, 22.5 * 3, 
		22.5 * 4, 22.5 * 5, 22.5 * 6, 22.5 * 7, 
		22.5 * 8, 22.5 * 9, 22.5 * 10, 22.5 * 11, 
		22.5 * 12, 22.5 * 13, 22.5 * 14, 22.5 * 15
	}
	correspX := [16]float32{ 
		0, 0.38268343, 0.70710678, 0.92387953, 
		1, 0.92387953, 0.70710678, 0.38268343, 
		0, -0.38268343, -0.70710678, -0.92387953, 
		-1, -0.92387953, -0.70710678, -0.38268343
	}
	correspY := [16]float32{ 
		1, 0.92387953, 0.70710678, 0.38268343,
		0, -0.38268343, -0.70710678, -0.92387953,
		-1, -0.92387953, -0.70710678, -0.38268343,
		0, 0.38268343, 0.70710678, 0.92387953
	}

	reading, err := winddirPin.Read()

	for i := 0; i < 16; i++ {
		if ( (dirValues[i] + 4) >= reading) && (reading >= (dirValues[i] - 4)) {
			winddirSum.x += correspX[i]
			winddirSum.y += correspY[i]
			return
		}
	}
	log.Fatal("couldn't find a corresponding direction")
}

func finalWindDir() (degree float32, dirName string) {
	winddirSum.x /= sleepCycles
	winddirSum.y /= sleepCycles

	
}



func zeroValues() {
	winddirSum.x = 0
	winddirSum.y = 0 

	temperatureSum = 0
	pressureSum = 0
	humiditySum = 0

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
		zeroValues()
		for i := 0; i < 60; i++ {
			
		}
	}


}
