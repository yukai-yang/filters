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
	data    *mults.MulTS // SetData
	from    int          // SetFrame
	to      int          // SetFrame
	nlatent int          // Init
	nvar    int          // Init
	nsample int          // Init
	parF    mat.Matrix
	parB    mat.Matrix
	parH    mat.Matrix
	parA    mat.Matrix
	parQ    mat.Matrix
	parR    mat.Matrix
	mz      mat.Matrix
	mu      mat.Matrix
	mx      mat.Matrix
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

	if obj.parA == nil {
		return errors.New("matrix A missing")
	}

	obj.nlatent, _ = obj.parF.Dims()

	var err error
	if obj.mz, err = obj.data.DepVars(obj.from, obj.to); err != nil {
		return errors.New("no dependent variable")
	}
	obj.nsample, obj.nvar = obj.mz.Dims()

	if obj.mu, err = obj.data.IndepVars(obj.from, obj.to); err != nil {
		return errors.New("no independent variable")
	}

	if obj.parQ == nil {
		obj.parQ = mat.NewDiagDense(obj.nlatent, filters.Repeat(1, obj.nlatent))
	}

	if obj.parR == nil {
		obj.parR = mat.NewDiagDense(obj.nvar, filters.Repeat(1, obj.nvar))
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

// SetPar sets the parameter matrices
func (obj *Kalman) SetPar(name string, par []float64, dim0, dim1 int) error {
	switch name {
	case "F":
		obj.parF = mat.NewDense(dim1, dim0, par).T()
	case "B":
		obj.parB = mat.NewDense(dim1, dim0, par).T()
	case "H":
		obj.parH = mat.NewDense(dim1, dim0, par).T()
	case "A":
		obj.parA = mat.NewDense(dim1, dim0, par).T()
	default:
		return errors.New("invalid parameter matrix name")
	}
	return nil
}
