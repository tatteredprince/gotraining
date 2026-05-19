package main

import (
	"errors"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// requestTimeout executes request with timeout.
func requestTimeout(request func(), timeout float64) (float64, error) {
	done := make(chan float64)
	go func() {
		start := time.Now()
		request()
		done <- time.Since(start).Seconds()
	}()
	ticker := time.Tick(500 * time.Millisecond)
	timer := time.After(time.Duration(timeout) * time.Second)
	for {
		select {
		case elapsed := <-done:
			fmt.Printf("elapsed %f seconds", elapsed)
			return elapsed, nil
		case <-timer:
			return -1, errors.New("timeout")
		case <-ticker:
			fmt.Println("request handling")
		}
	}
}

func requestTimeoutTestHelper(t *testing.T, request func(), timeout float64, success bool) {
	t.Helper()
	if success {
		t.Logf("expect request to take less than %f seconds", timeout)
	} else {
		t.Logf("expect request to last more than %f seconds", timeout)
	}
	if _, err := requestTimeout(request, timeout); success && err != nil || !success && err == nil {
		t.Fatal()
	}
}

func TestRequestTimeout(t *testing.T) {
	t.Run("Long lasting request", func(t *testing.T) {
		requestTimeoutTestHelper(t, func() {
			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			time.Sleep(time.Duration(1+r.Intn(3)) * time.Second)
		}, 1.0, false)
	})
	t.Run("Fast request", func(t *testing.T) {
		requestTimeoutTestHelper(t, func() {}, 1.0, true)
	})
}
