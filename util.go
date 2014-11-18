package hearthscience

import "math"

func round(n float32) float32 {
	f64 := float64(n)
	if f64 < 0 {
		return float32(math.Ceil(f64 - 0.5))
	}
	return float32(math.Floor(f64 + 0.5))
}
