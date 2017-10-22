package climit

import (
	"sync"
	"sync/atomic"
	"testing"
)

func TestNew(t *testing.T) {
	const limit = 3

	if New(limit) == nil {
		t.Fatal("New returns a nil channel")
	}
}

func BenchmarkNew2(b *testing.B) {
	const limit = 2
	var cl ConcurrencyLimit

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cl = New(limit)
	}
	b.StopTimer()

	if cl == nil {
	}
}

func TestExec(t *testing.T) {
	const limit = 5
	const instances = 10

	counter := new(uint64)
	wg := new(sync.WaitGroup)

	cl := New(limit)
	for i := 0; i < instances; i++ {
		if i < limit {
			wg.Add(1)
		}
		go cl.Exec(func() {
			atomic.AddUint64(counter, 1)
			wg.Done()
			select {}
		})
	}
	wg.Wait()

	runningInstancesCount := *counter
	if runningInstancesCount != limit {
		t.Fatalf("Too many instances are running. Expected: %d, Actual: %d", limit, runningInstancesCount)
	}
}

func BenchmarkExec2(b *testing.B) {
	const limit = 2

	b.StopTimer()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		cl := New(limit)
		for instances := 0; instances < limit; instances++ {
			b.StartTimer()
			cl.Exec(func() {})
			b.StopTimer()
		}
	}
}
