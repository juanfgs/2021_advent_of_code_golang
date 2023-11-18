package sonar_sweep

func Measurements(measurements []int) int {
	var previousMeasure int = 0
	increments := 0
	for i, v := range measurements {
		if previousMeasure < v && i != 0 {
			increments++
		}

		previousMeasure = v
	}

	return increments
}
