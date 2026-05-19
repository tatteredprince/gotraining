package main

import (
	"slices"
	"sync"
	"testing"
)

// fanOut returns output channels filled with values from input.
func fanOut(input <-chan int, num int) []chan int {
	output := make([]chan int, num)
	for i := range num {
		output[i] = make(chan int)
	}
	wg := sync.WaitGroup{}
	wg.Add(num)
	for i := range num {
		go func() {
			defer wg.Done()
			for val := range input {
				output[i] <- val
			}
		}()
	}
	go func() {
		wg.Wait()
		for i := range num {
			close(output[i])
		}
	}()
	return output
}

func TestFanOut(t *testing.T) {
	t.Run("Five workers on ten jobs", func(t *testing.T) {
		// Fill input with jobs
		jobs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		input := make(chan int, len(jobs))
		for _, j := range jobs {
			input <- j
		}
		close(input)

		// Fan out jobs
		numWorkers := 5
		chans := fanOut(input, numWorkers)

		// Gather finished jobs into output
		output := make(chan int)
		wg := sync.WaitGroup{}
		wg.Add(numWorkers)
		for i := range numWorkers {
			go func() {
				defer wg.Done()
				for val := range chans[i] {
					output <- val
				}
			}()
		}
		go func() {
			wg.Wait()
			close(output)
		}()

		// Check finished jobs in output
		finJobs := make([]int, 0, len(jobs))
		for j := range output {
			finJobs = append(finJobs, j)
		}
		slices.Sort(finJobs)
		t.Logf("got finished jobs %v", finJobs)
		if !slices.Equal(finJobs, jobs) {
			t.Fatal()
		}
	})
}
