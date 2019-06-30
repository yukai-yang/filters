package filters

/* the common interface */

// Filter defines the common methods of the filters
type Filter interface {
	Filtering() error
	Smoothing() error
}
