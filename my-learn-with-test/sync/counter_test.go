package sync

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := NewCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()
		assertCounter(t, counter, 3)
	})

	t.Run("increment async", func(t *testing.T) {
		wantedCount := 100
		counter := NewCounter()

		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			go func(_wg *sync.WaitGroup) {
				counter.Inc()
				_wg.Done()
			}(&wg)
		}
		wg.Wait()

		assertCounter(t, counter, wantedCount)
	})
}

func assertCounter(t *testing.T, counter *Counter, v int) {
	// this is to indicate that this method is helper, so if the test breaks, the error would be shown
	// at the caller
	t.Helper()
	if counter.Value() != v {
		t.Errorf("got %d, want %d", counter.Value(), v)
	}
}
