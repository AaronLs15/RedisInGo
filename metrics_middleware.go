package main

import (
	"fmt"
	"redisgo/store"
	"time"
)

type MetricsMiddleware struct {
	inner store.Storer

	getCalls    int
	setCalls    int
	deleteCalls int
	lenCalls    int

	getMisses       int
	totalGetLatency time.Duration
	totalSetLatency time.Duration
}

func NewMetricsMiddleware(inner store.Storer) *MetricsMiddleware {
	return &MetricsMiddleware{inner: inner}
}

func (m *MetricsMiddleware) Get(Key string) (string, error) {
	start := time.Now()
	val, err := m.inner.Get(Key)

	if err != nil {
		m.getMisses++
	}

	m.getCalls++

	m.totalGetLatency += time.Since(start)

	return val, err
}

func (m *MetricsMiddleware) Set(key string, vamue string) error {
	start := time.Now()

	err := m.inner.Set(key, vamue)
	m.setCalls++

	m.totalSetLatency += time.Since(start)

	return err
}

func (m *MetricsMiddleware) Delete(key string) {
	m.inner.Delete(key)
	m.deleteCalls++
}

func (m *MetricsMiddleware) Len() int {
	got := m.inner.Len()
	m.lenCalls++
	return got
}

func (m *MetricsMiddleware) Keys() []string {
	got := m.inner.Keys()
	return got
}

func (m *MetricsMiddleware) GetMetrics() {

	fmt.Printf("GET calls: %d (misses %d)\n", m.getCalls, m.getMisses)
	fmt.Printf("SET calls: %d\n", m.setCalls)
	fmt.Printf("DEL calls: %d\n", m.deleteCalls)
	fmt.Printf("LEN calls: %d\n", m.lenCalls)

	if m.getCalls > 0 {
		avg := m.totalGetLatency / time.Duration(m.getCalls)
		fmt.Printf("AVG GET latency: %s\n", avg)
	}
	if m.setCalls > 0 {
		avg := m.totalGetLatency / time.Duration(m.setCalls)
		fmt.Printf("AVG SET latency: %s\n", avg)
	}
}
