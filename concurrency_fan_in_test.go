package main

import (
	"slices"
	"sync"
	"testing"
)

// fanIn returns channel that multiplexes values from input channels.
func fanIn(inputs ...chan int) <-chan int {
	output := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(len(inputs))
	for _, ch := range inputs {
		go func() {
			defer wg.Done()
			for value := range ch {
				output <- value
			}
		}()
	}
	go func() {
		wg.Wait()
		close(output)
	}()
	return output
}

func TestFanIn(t *testing.T) {
	t.Run("Count to ten", func(t *testing.T) {
		// Fill inputs
		expect := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		inputs := make([]chan int, 0, len(expect))
		for range expect {
			inputs = append(inputs, make(chan int))
		}
		for i, val := range expect {
			go func() {
				inputs[i] <- val
				close(inputs[i])
			}()
		}

		// Read from multiplexed output
		got := make([]int, 0, len(expect))
		for num := range fanIn(inputs...) {
			got = append(got, num)
		}

		// Check received values
		slices.Sort(got)
		slices.Sort(expect)
		if !slices.Equal(got, expect) {
			t.Fatalf("expected %v but got %v", expect, got)
		}
	})
}
