package main

import (
	"flag"
	"fmt"
	"sync"
)

// fetcher is abstraction for sending network request, so we are able to replace implementations with test one
// and make tests independent of network availability
type fetcher interface {
	Fetch(url string) (urlAndHash urlAndHash)
}

type runner struct {
	parallel int
	taskCh   chan string
	respCh   chan urlAndHash
	wg       *sync.WaitGroup
	fetcher  fetcher
}

type urlAndHash struct {
	url  string
	hash string
}

func newRunner(parallel int, fetcher fetcher) *runner {
	wg := &sync.WaitGroup{}
	wg.Add(parallel)

	return &runner{
		parallel: parallel,
		taskCh:   make(chan string),
		respCh:   make(chan urlAndHash),
		wg:       wg,
		fetcher:  fetcher,
	}
}

func (r *runner) run(urls []string) []urlAndHash {
	// start workers of parallel parameter count
	for i := 0; i < r.parallel; i++ {
		go r.worker()
	}

	resultsCompletelyReadCh := make(chan struct{})
	var result []urlAndHash
	// start reading responses from workers and add them to result slice
	go func() {
		for resp := range r.respCh {
			result = append(result, resp)
		}
		close(resultsCompletelyReadCh)
	}()

	// send tasks to task channel which will be processed by the workers
	for _, url := range urls {
		r.taskCh <- url
	}
	// close taskCh so workers understand that they should stop and return
	close(r.taskCh)

	// wait until all workers returned
	r.wg.Wait()
	close(r.respCh)
	// wait until all results read
	<-resultsCompletelyReadCh

	return result
}

func (r *runner) worker() {
	defer r.wg.Done()

	for taskURL := range r.taskCh {
		r.respCh <- r.fetcher.Fetch(taskURL)
	}
}

func main() {
	parallelPtr := flag.Int("parallel", 10, "parallel")
	flag.Parse()
	tail := flag.Args()

	r := newRunner(*parallelPtr, newHTTPFetcher())
	results := r.run(tail)

	for _, result := range results {
		fmt.Println(result.url, result.hash)
	}
}
