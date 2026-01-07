package main

import "testing"

func expectFailure(t *testing.T, target int, complements []int) {
	t.Helper()
	t.Logf("expecting failed search for two complenents in %d for value %d", complements, target)
	if hasTwoComplements(target, complements) {
		t.Fatalf("%d has complement in %d", target, complements)
	}
}

func expectSuccess(t *testing.T, target int, complements []int) {
	t.Helper()
	t.Logf("expecting successful search for two complements in %d for value %d", complements, target)
	if !hasTwoComplements(target, complements) {
		t.Fatalf("%d has no complement in %d", target, complements)
	}
}

func TestTwoComplements(t *testing.T) {
	t.Run("Binary set", func(t *testing.T) { expectSuccess(t, 1, []int{0, 1}) })
	t.Run("Continious binary set", func(t *testing.T) { expectFailure(t, 9, []int{1, 0, 1, 0, 1, 0, 1}) })
	t.Run("Empty set", func(t *testing.T) { expectFailure(t, 5, []int{}) })
	t.Run("Zeroes set", func(t *testing.T) { expectSuccess(t, 0, []int{0, 0, 0, 0, 0}) })
	t.Run("No complement", func(t *testing.T) { expectFailure(t, 3, []int{3, 4, 5, 6, 7}) })
}
