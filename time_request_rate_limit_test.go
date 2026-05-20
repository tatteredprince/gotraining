package main

import (
	"fmt"
	"testing"
	"time"
)

// requestRateLimit executes request with rate limiting.
func requestRateLimit(request func(int), input <-chan int, rateLimit int) {
	ticker := time.Tick(time.Second / time.Duration(rateLimit))
	for j := range input {
		<-ticker
		request(j)
	}
}

func TestRateLimit(t *testing.T) {
	t.Run("Several requests with small rate limit", func(t *testing.T) {
		// Fill input with numbers
		requests := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		numRequests := len(requests)
		input := make(chan int, numRequests)
		for _, n := range requests {
			input <- n
		}
		close(input)

		// Execute request within rate limit
		rateLimit := 1
		start := time.Now()
		requestRateLimit(
			func(n int) {
				fmt.Printf("processing request with id %d\n", n)
			}, input, rateLimit,
		)
		elapsed := time.Since(start)

		// Check elapsed time
		expect := time.Second * time.Duration(numRequests) / time.Duration(rateLimit)
		t.Logf("expect request with rate limit of %d/s to take more than %s", rateLimit, expect.String())
		t.Logf("request took %s", elapsed.String())
		if elapsed < expect {
			t.Fatal()
		}
	})
}
