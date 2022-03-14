package main

import (
	"sort"
	"strconv"
	"sync/atomic"
	"testing"
)

func TestRunner_Run(t *testing.T) {
	tests := []struct {
		name     string
		parallel int
		urls     []string
	}{
		{
			name:     "single_thread",
			parallel: 1,
			urls:     []string{"hello", "there", "world"},
		},
		{
			name:     "two_thread",
			parallel: 2,
			urls:     []string{"hello", "there", "world", "peace", "programming", "and", "so", "on"},
		},
		{
			name:     "more_threads_then_tasks",
			parallel: 100,
			urls:     []string{"hello", "there", "world", "peace", "programming", "and", "so", "on"},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := newRunner(test.parallel, &mockFetcher{}).run(test.urls)
			if len(result) != len(test.urls) {
				t.Error("count of results missmatch")
			}
			sort.Slice(result, func(i, j int) bool {
				return result[i].hash < result[j].hash
			})
			for i, resultItem := range result {
				if resultItem.hash != "hash"+strconv.FormatInt(int64(i+1), 10) {
					t.Error("invalid result")
				}
			}
		})
	}
}

type mockFetcher struct {
	counter int32
}

func (m *mockFetcher) Fetch(url string) urlAndHash {
	newCounter := atomic.AddInt32(&m.counter, 1)
	return urlAndHash{
		url:  url,
		hash: "hash" + strconv.FormatInt(int64(newCounter), 10),
	}
}
