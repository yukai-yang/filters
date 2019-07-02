package kalman

import (
	"errors"

	"github.com/yukai-yang/mults"
	"gonum.org/v1/gonum/mat"
)

// Kalman defines the structure of the Kalman filter
type Kalman struct {
	// contains filtered or unexported fields
	data *mults.MulTS // SetData
	from int          // SetFrame
	to   int          // SetFrame
	parF mat.Matrix
	parB mat.Matrix
	parH mat.Matrix
	parQ mat.Matrix
	parR mat.Matrix
}

/* functions for the Filter interface */

// Init does the initialization if any
func (obj *Kalman) Init() error {
	return nil
}

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
	var mfrom, mto = obj.data.PossibleFrame()
	if from < mfrom || from >= to || to > mto {
		return errors.New("invalid from or to")
	}
	obj.from = from
	obj.to = to
	return nil
}
