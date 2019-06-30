// Package kalman provides implementations of the Kalman filtering algorithm
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
