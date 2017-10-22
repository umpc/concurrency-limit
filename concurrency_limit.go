package climit

// ConcurrencyLimit is a type used for limiting concurrently-running instances of a function or group of functions.
type ConcurrencyLimit chan struct{}

// New is a convienience wrapper for: make(ConcurrencyLimit, limit)
func New(limit int) ConcurrencyLimit {
	return make(ConcurrencyLimit, limit)
}

// Exec runs a given function when the number of active goroutines falls below the limit set in New.
func (cl ConcurrencyLimit) Exec(fn func()) {
	cl <- struct{}{}
	fn()
	<-cl
}
