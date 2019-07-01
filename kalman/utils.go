package kalman

import "github.com/yukai-yang/mults"

/* utility functions */

// NewKalman initializes the Kalman filter
// and returns the address of an object of Kalman struct
func NewKalman(data *mults.MulTS) *Kalman {
	var kal = &Kalman{}
	kal.data = data
	return kal
}
