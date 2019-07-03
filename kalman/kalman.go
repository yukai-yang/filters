package kalman

import (
	"errors"

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
	nindep  int          // Init
	parF    mat.Matrix   // SetPar
	parB    mat.Matrix   // SetPar
	parH    mat.Matrix   // SetPar
	parA    mat.Matrix   // SetPar
	parQ    mat.Matrix   // SetPar, Init
	parR    mat.Matrix   // SetPar, Init
	mz      mat.Matrix   // obs, Init
	mu      mat.Matrix   // exp obs, Init
	mx      mat.Matrix   // pred latent, Update
	mxx     mat.Matrix   // updt latent, Update
	aP      []mat.Matrix // pred cov Q
	aPP     []mat.Matrix // updt cov Q
	mv      mat.Matrix   // pred noise obs (y)
	mvv     mat.Matrix   // updt noise obs (y)
	aS      []mat.Matrix
	aK      []mat.Matrix
	inix    mat.Matrix
	iniP    mat.Matrix
}

/* functions for the Filter interface */

// Init does the initialization
// basically it checks the parameters of the model
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

	if tmp1, _ := obj.parF.Dims(); tmp1 != obj.nlatent {
		return errors.New("invalid F dimension")
	}

	var err error
	if obj.mz, err = obj.data.DepVars(obj.from, obj.to); err != nil {
		return err
	}
	obj.nsample, obj.nvar = obj.mz.Dims()

	if tmp1, tmp2 := obj.parH.Dims(); tmp1 != obj.nvar || tmp2 != obj.nlatent {
		return errors.New("invalid H dimension")
	}

	if obj.mu, err = obj.data.IndepVars(obj.from, obj.to); err != nil {
		return err
	}
	_, obj.nindep = obj.mu.Dims()

	if tmp1, tmp2 := obj.parA.Dims(); tmp1 != obj.nvar || tmp2 != obj.nindep {
		return errors.New("invalid A dimension")
	}

	if tmp1, tmp2 := obj.parB.Dims(); tmp1 != obj.nlatent || tmp2 != obj.nindep {
		return errors.New("invalid B dimension")
	}

	if obj.parQ == nil {
		obj.parQ = mat.NewDiagDense(obj.nlatent, repeat(1, obj.nlatent))
	} else {
		if tmp1, _ := obj.parQ.Dims(); tmp1 != obj.nlatent {
			return errors.New("invalid Q dimension")
		}
	}

	if obj.parR == nil {
		obj.parR = mat.NewDiagDense(obj.nvar, repeat(1, obj.nvar))
	} else {
		if tmp1, _ := obj.parR.Dims(); tmp1 != obj.nvar {
			return errors.New("invalid R dimension")
		}
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
		if dim0 != dim1 || dim0 < 1 || dim1 < 1 || dim0*dim1 != len(par) {
			return errors.New("invalid parameter matrix dimensions")
		}
		obj.parF = mat.NewDense(dim1, dim0, par).T()
	case "B":
		if dim0 < 1 || dim1 < 1 || dim0*dim1 != len(par) {
			return errors.New("invalid parameter matrix dimensions")
		}
		obj.parB = mat.NewDense(dim1, dim0, par).T()
	case "H":
		if dim0 < 1 || dim1 < 1 || dim0*dim1 != len(par) {
			return errors.New("invalid parameter matrix dimensions")
		}
		obj.parH = mat.NewDense(dim1, dim0, par).T()
	case "A":
		if dim0 < 1 || dim1 < 1 || dim0*dim1 != len(par) {
			return errors.New("invalid parameter matrix dimensions")
		}
		obj.parA = mat.NewDense(dim1, dim0, par).T()
	case "Q":
		// only dim0 will be used
		if dim0 < 1 || dim0*dim0 != len(par) {
			return errors.New("invalid parameter matrix dimensions")
		}
		obj.parQ = mat.NewSymDense(dim0, transposefloats(par, dim0, dim0))
	case "R":
		// only dim0 will be used
		if dim0 < 1 || dim0*dim0 != len(par) {
			return errors.New("invalid parameter matrix dimensions")
		}
		obj.parR = mat.NewSymDense(dim0, transposefloats(par, dim0, dim0))
	default:
		return errors.New("invalid parameter matrix name")
	}
	return nil
}
