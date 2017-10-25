package math

func Average( xs []float64) float64 {
	var total float64 = 0
	for _, x:= range xs {
		total += x
	}
	return total / float64(len(xs))
}