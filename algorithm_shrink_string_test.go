package main

import "testing"

// shrinkString substitutes duplicate letters in string with counters.
func shrinkString(str string) string {
	shrinked := ""
	return shrinked
}

func shrinkStringTestHelper(t *testing.T, str, expect string) {
	t.Helper()
	t.Logf("shrinking '%s', expecting '%s'", str, expect)
	got := shrinkString(str)
	t.Logf("got shrinked '%s'", got)
	if got != expect {
		t.Fail()
	}
}

func TestShrinkString(t *testing.T) {
	t.Run("Empty string", func(t *testing.T) { shrinkStringTestHelper(t, "", "") })
	t.Run("Short string", func(t *testing.T) { shrinkStringTestHelper(t, "ABC", "ABC") })
	t.Run("Long string", func(t *testing.T) { shrinkStringTestHelper(t, "AAAABBBCCXYZ", "A4A3C2XYZ") })
}
