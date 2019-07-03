package kalman

import (
	"github.com/yukai-yang/mults"
	"gonum.org/v1/gonum/mat"
)

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

// SymCrossProd performs a product between x', a and x,
//  s = x' * a * x * b
func symCrossProd(a mat.Symmetric, x mat.Matrix, b float64) *mat.SymDense {
	if x == nil {
		panic(mat.ErrZeroLength)
	}
	n, m := x.Dims()

	if a != nil {
		if a.Symmetric() != n {
			panic(mat.ErrShape)
		}
	} else {
		a = mat.NewDiagDense(n, repeat(1, n))
	}

	s := mat.NewSymDense(m, nil)
	xx := mat.DenseCopyOf(x)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			s.RankTwo(s, b*a.At(i, j), xx.RowView(i), xx.RowView(j))

		}
	}

	return s
}
