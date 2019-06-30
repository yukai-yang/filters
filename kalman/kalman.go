package kalman

import "gonum.org/v1/gonum/mat"

// Kalman defines the structure of the Kalman filter
type Kalman struct {
	// contains filtered or unexported fields
	obs  *mat.Dense
	parF *mat.Dense
	parB *mat.Dense
	parH *mat.Dense
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
func (obj *Kalman) SetObs(obs []float64) {

}
