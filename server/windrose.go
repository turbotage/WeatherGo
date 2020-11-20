package server

func getDirectionName(degree float32) string {
	dirList := [16]string{"N", "NNE", "NE", "ENE", "E", "ESE", "SE", "SSE", "S", "SSW", "SW", "WSW", "W", "WNW", "NW", "NNW"}
	if (degree > 348.75) || ((degree >= 0.0) && (degree < 12.25)) {
		return "N"
	}
	for i := 1; i < 16; i++ {
		if ((22.5*float32(i) - 12.25) <= degree) && (degree < (22.5*float32(i) + 12.25)) {
			return dirList[i]
		}
	}
	return "N"
}

func getDirectionNumber(degree float32) int {
	if (degree > 348.75) || ((degree >= 0.0) && (degree < 12.25)) {
		return 0
	}
	for i := 1; i < 16; i++ {
		if ((22.5*float32(i) - 12.25) <= degree) && (degree < (22.5*float32(i) + 12.25)) {
			return i
		}
	}
	return 0
}
