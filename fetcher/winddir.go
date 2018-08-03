package fetcher

type vec2 struct {
	x float32
	y float32
}

var winddirPin AnalogPin
var winddirSum vec2

func zeroWindDir(){
	winddirSum.x = 0
	winddirSum.y = 0 
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

	dirNames := [16]string{"N", "NNE", "NE", "ENE", "E", "ESE", "SE", "SSE", "S", "SSW", "SW", "WSW", "W", "WNW", "" }

	if (winddirSum.x >= 0) && (winddirSum.y >= 0) {
		degree = 90 - math.Atan2(winddirSum.y, winddirSum.x)
	} else if winddirSum.y < 0 {
		degree = (-1 * math.Atan2(winddirSum.y, winddirSum.x)) + 90
	} else if (winddirSum.x < 0) && (winddirSum.y >= 0) {
		degree = 460 - math.Atan2(winddirSum.y, winddirSum.x);
	}

	for i := 0; i < 16; ++i {
		if ( degree > i * 22.5 ) && ( degree < (i + 1) * 22.5 ) {
			dirName = dirNames[i]
		}
	}

}

