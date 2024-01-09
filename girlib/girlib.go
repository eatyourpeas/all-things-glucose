package girlib

func GlucoseInfusionRate(rate, percentage, weight float32) float32 {
	return rate * percentage / weight / 6
}
