package kalman

import (
	"errors"

	"gonum.org/v1/gonum/mat"
)

// Kalman defines the structure of the Kalman filter
type Kalman struct {
	// contains filtered or unexported fields
	obs   *mat.Dense
	iTT   int
	begin int
	end   int
	parF  *mat.Dense
	parB  *mat.Dense
	parH  *mat.Dense
	parQ  *mat.Dense
	parR  *mat.Dense
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

// SetObs sets the observations
func (obj *Kalman) SetObs(data []float64, nvar int) error {
	obj.iTT = len(data) / nvar
	if obj.iTT*nvar != len(data) {
		return errors.New("dimensions do not fit")
	}
	obj.obs = mat.NewDense(obj.iTT, nvar, data)
	return nil
}

// SetFrame sets the begin and end of the time series
func (obj *Kalman) SetFrame(begin, end int) error {
	if begin < 0 || begin > end || end >= obj.iTT {
		return errors.New("begin or end are wrong")
	}
	obj.begin = begin
	obj.end = end
	return nil
}
