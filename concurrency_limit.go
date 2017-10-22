package climit

// ConcurrencyLimit is a type used for limiting concurrently-running instances of a function or group of functions.
type ConcurrencyLimit chan struct{}

// New accepts a limit value greater than zero and initializes a ConcurrencyLimit value
func New(limit int) ConcurrencyLimit {
	if limit < 1 {
		limit = 1
	}
	return make(ConcurrencyLimit, limit)
}

// Exec runs a given function when the number of active goroutines falls below the limit set in New.
func (cl ConcurrencyLimit) Exec(fn func()) {
	cl <- struct{}{}
	fn()
	<-cl
}
