package kalman

import "github.com/yukai-yang/mults"

/* utility functions */

// NewKalman initializes the Kalman filter
// and returns the address of an object of Kalman struct
func NewKalman(data *mults.MulTS) *Kalman {
	var kal = &Kalman{}
	kal.data = data
	return kal
}

func repeat(v float64, ntimes int) []float64 {
	var ret = make([]float64, ntimes)
	for i := 0; i < ntimes; i++ {
		ret[i] = v
	}
	return ret
}

func transposefloats(vals []float64, dim0, dim1 int) []float64 {
	var ret = make([]float64, len(vals))
	for i := 0; i < dim0; i++ {
		for j := 0; j < dim1; j++ {
			ret[j+i*dim1] = vals[i+j*dim0]
		}
	}
	return ret
}
