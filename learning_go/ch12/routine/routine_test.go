package routine

import (
	"sync"
	"testing"
)

func TestSummationWithAtomicOperations(t *testing.T) {
	t.Parallel()
	for i := 0; i < 10000; i++ {
		var wg sync.WaitGroup
		var result int32
		wg.Add(4)
		go SummationWithAtomicOperations(IterationParams{Start: 34, Stop: 65, Step: 1}, &wg, &result)
		go SummationWithAtomicOperations(IterationParams{Start: 66, Stop: 88, Step: 1}, &wg, &result)
		go SummationWithAtomicOperations(IterationParams{Start: 89, Stop: 100, Step: 1}, &wg, &result)
		go SummationWithAtomicOperations(IterationParams{Start: 0, Stop: 33, Step: 1}, &wg, &result)
		wg.Wait()
		if result != 5050 {
			t.Errorf("expected 5050, got %d", result)
		}
	}
}
