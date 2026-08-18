package prometheus

import (
	"sync"
)

type Labels map[string]string

type MetricVec struct {
	mu          sync.RWMutex
	metricMap   map[string]interface{}
}

func NewMetricVec() *MetricVec {
	return &MetricVec{
		metricMap: make(map[string]interface{}),
	}
}

func (v *MetricVec) GetMetricWith(labels Labels) (interface{}, error) {
	v.mu.Lock()
	defer v.mu.Unlock()

	// Defensive copy of labels to prevent external mutation
	key := ""
	for k, val := range labels {
		key += k + "=" + val + ","
	}

	if m, ok := v.metricMap[key]; ok {
		return m, nil
	}

	// Logic to create new metric would go here
	return nil, nil
}