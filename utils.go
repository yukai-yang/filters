package filters

/* utility functions */

func rep(v float64, ntimes int) []float64 {
	var ret = make([]float64, ntimes)
	for i := 0; i < ntimes; i++ {
		ret[i] = v
	}
	return ret
}
