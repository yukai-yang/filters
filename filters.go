package filters

import "github.com/yukai-yang/mults"

/* the common interface */

// Filter defines the common methods of the filters
type Filter interface {
	Init() error
	Filtering() error
	Smoothing() error
	// data and model
	SetData(*mults.MulTS)
	SetFrame(int, int) error
	// name of par, values of par, and two dimensions
	SetPar(string, []float64, int, int) error
}
