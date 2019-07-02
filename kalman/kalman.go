package kalman

import (
	"errors"

	"github.com/yukai-yang/filters"
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
	if obj.data == nil {
		return errors.New("no data")
	}

	if obj.from == 0 && obj.to == 0 {
		obj.from, obj.to = obj.data.PossibleFrame()
	}

	if obj.parF == nil {
		return errors.New("matrix F missing")
	}

	if obj.parB == nil {
		return errors.New("matrix B missing")
	}

	if obj.parH == nil {
		return errors.New("matrix H missing")
	}

	var nlatent, _ = obj.parF.Dims()

	if obj.parQ == nil {
		obj.parQ = mat.NewDiagDense(nlatent, filters.Repeat(1, nlatent))
	}

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
