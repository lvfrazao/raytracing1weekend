package vec3

func clamp(x, min, max float64) float64 {
	// Clamps the x value between min and max
	if x < min {
		return min
	} else if x > max {
		return max
	}
	return x
}
