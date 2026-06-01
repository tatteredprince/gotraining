package main

import "testing"

// anagramsSearch determines if string contains substring equivalent to anagram.
func anagramsSearch(str, anagram string) bool {
	return false
}

func anagramsSearchTestHelper(t *testing.T, str, anagram string, expect bool) {
	t.Helper()
	t.Logf("searching anagram '%s' in string '%s'", anagram, str)
	got := anagramsSearch(str, anagram)
	if got != expect {
		t.Fatal()
	}
}

func TestAnagramsSearch(t *testing.T) {
	t.Run("Simple string", func(t *testing.T) { anagramsSearchTestHelper(t, "reebok", "eeb", true) })
	t.Run("Complicated string", func(t *testing.T) { anagramsSearchTestHelper(t, "abracadabra", "aad", true) })
	t.Run("Funny string", func(t *testing.T) { anagramsSearchTestHelper(t, "hullabaloo", "lull", false) })
	t.Run("Nonoverlapping letters", func(t *testing.T) { anagramsSearchTestHelper(t, "achive", "xyz", false) })
}
