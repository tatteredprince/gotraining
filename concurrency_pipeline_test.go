package main

import (
	"slices"
	"sync"
	"testing"
)

// pipelineFilterOddNumbers cuts off odd numbers from input.
func pipelineFilterOddNumbers(input <-chan int) <-chan int {
	output := make(chan int)
	go func() {
		defer close(output)
		for v := range input {
			if v%2 == 0 {
				output <- v
			}
		}
	}()
	return output
}

// pipelineSquaring squares numbers from input.
func pipelineSquaring(input <-chan int) <-chan int {
	output := make(chan int)
	go func() {
		defer close(output)
		for v := range input {
			output <- v * v
		}
	}()
	return output
}

// pipelineDoubling doubles numbers from input.
func pipelineDoubling(input <-chan int) <-chan int {
	output := make(chan int)
	go func() {
		defer close(output)
		for v := range input {
			output <- v + v
		}
	}()
	return output
}

func TestPipeline(t *testing.T) {
	t.Run("Pipeline ten numbers", func(t *testing.T) {
		// Fill input with numbers.
		numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		input := make(chan int, len(numbers))
		for _, n := range numbers {
			input <- n
		}
		close(input)

		// Run pipeline.
		expect := []int{0, 2 * 2 * 2, 4 * 4 * 2, 6 * 6 * 2, 8 * 8 * 2}
		got := make([]int, 0, len(expect))
		output := pipelineDoubling(pipelineSquaring(pipelineFilterOddNumbers(input)))
		wg := sync.WaitGroup{}
		wg.Add(1)
		go func() {
			defer wg.Done()
			for n := range output {
				got = append(got, n)
			}
		}()
		wg.Wait()

		// Check pipeline's output.
		slices.Sort(expect)
		slices.Sort(got)
		t.Logf("expect numbers %v but got %v", expect, got)
		if !slices.Equal(expect, got) {
			t.Fatal()
		}
	})
}
