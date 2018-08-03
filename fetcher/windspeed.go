package fetcher

var anemometerPinNum uint8
var anemometerPin InterruptPin
var anemometerCount uint32

var gustCounters[8]uint32
var highestGust float32

func setupWindSpeedAndGust() {
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

	anemometerPin.Watch(embd.EdgeFalling, func(){
		anemometerCount++

		//gust
		for i := 0; i < 8; ++i {
			gustCounters++
		}
	})
}

func zeroWindSpeedAndGust() {
	
}

func runGustCycle() {
	
	highestGust = 0;
	
	for i := 0; i < 30 * 5; ++i {
		for j := 0; i < 8; ++i {
			gustCounters[j] = 0
			time.Sleep(2000 * time.Millisecond / 8)
		}
		for j := 0; i < 8; ++i {
			thisGust := gustCounters[j] * 2.40114125 / 3.6
			if thisGust > highestGust {
				highestGust = thisGust
			}
			time.Sleep(2000 * time.Millisecond / 8)
		}
	}
}

