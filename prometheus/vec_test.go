package prometheus

import (
	"sync"
	"testing"
)

func TestMetricVec_ConcurrentWith_Race(t *testing.T) {
	vec := NewMetricVec()

	var wg sync.WaitGroup
	concurrency := 100
	iterations := 1000

	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < iterations; j++ {
				_, _ = vec.GetMetricWith(Labels{})
			}
		}()
	}
	wg.Wait()
}