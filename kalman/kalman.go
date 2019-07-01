package kalman

import (
	"github.com/yukai-yang/mults"
	"gonum.org/v1/gonum/mat"
)

// Kalman defines the structure of the Kalman filter
type Kalman struct {
	// contains filtered or unexported fields
	data *mults.MulTS
	from int
	to   int
	parF *mat.Dense
	parB *mat.Dense
	parH *mat.Dense
	parQ *mat.Dense
	parR *mat.Dense
}

/* functions for the Filter interface */

// Filtering does the Kalman filtering
func (obj *Kalman) Filtering() error {

	return nil
}

// Smoothing does the disturbance smoothing
func (obj *Kalman) Smoothing() error {

	return nil
}

/* Kalman methods */

// SetData sets the observations
func (obj *Kalman) SetData(data *mults.MulTS) {
	obj.data = data
}

// SetFrame sets the from and to of the time series
func (obj *Kalman) SetFrame(from, to int) error {
	obj.from = from
	obj.to = to
	return nil
}
