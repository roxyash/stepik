// t = 6/w; w = sqrt(k/m); m = p*v
func M() float64 {
	return p * v
}

func W() float64 {
	return math.Sqrt(k/M())
}

func T() float64 {
	return 6/W()
}