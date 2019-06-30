package kalman

/* utility functions */

// NewKalman initializes the Kalman filter
// and returns the address of an object of Kalman struct
func NewKalman() *Kalman {
	var kal = &Kalman{}
	return kal
}
