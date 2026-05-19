package main

import (
	"slices"
	"testing"
)

// generateFibonacciNumbers returns channel that contains Fibonacci numbers.
func generateFibonacciNumbers(first, second, count int) <-chan int {
	out := make(chan int)
	go func() {
		for range count {
			first, second = second, first+second
			out <- second
		}
		close(out)
	}()
	return out
}

func generateFibonacciNumbersTestHelper(t *testing.T, first, second, count int, expect []int) {
	t.Helper()
	t.Logf("calculate %d fibonacci numbers after %d and %d", count, first, second)
	got := make([]int, 0, count)
	for num := range generateFibonacciNumbers(first, second, count) {
		got = append(got, num)
	}
	if !slices.Equal(got, expect) {
		t.Fatalf("expected %v but got %v", expect, got)
	}
}

func TestGenerateFibonacciNumbers(t *testing.T) {
	t.Run("Next", func(t *testing.T) { generateFibonacciNumbersTestHelper(t, 0, 1, 1, []int{1}) })
	t.Run("First five", func(t *testing.T) { generateFibonacciNumbersTestHelper(t, 0, 1, 5, []int{1, 2, 3, 5, 8}) })
	t.Run("Double digits", func(t *testing.T) { generateFibonacciNumbersTestHelper(t, 5, 8, 5, []int{13, 21, 34, 55, 89}) })
	t.Run("Triple digits", func(t *testing.T) {
		generateFibonacciNumbersTestHelper(t, 55, 89, 5, []int{144, 233, 377, 610, 987})
	})
	t.Run("Big numbers", func(t *testing.T) {
		generateFibonacciNumbersTestHelper(t, 832040, 1346269, 10, []int{2178309, 3524578, 5702887, 9227465, 14930352, 24157817, 39088169, 63245986, 102334155, 165580141})
	})
}
