// Package filters implements several filtering algorithms
package filters

import "gonum.org/v1/gonum/mat"

/* the interface */

// Filter defines the common methods of the filters
type Filter interface {
}

// Kalman defines the structure of the Kalman filter
type Kalman struct {
	// contains filtered or unexported fields
	obs  *mat.Dense
	parF *mat.Dense
	parB *mat.Dense
	parH *mat.Dense
}
