package projecteuler

import (
	"testing"
)

func TestAnswer(t *testing.T) {
	cases := []struct {
		digits     int
		palindrome int
	}{
		{2, 9009},
		{3, 906609},
	}

	for _, c := range cases {
		got := answer(c.digits)
		want := c.palindrome

		if got != want {
			t.Errorf("answer(%d) = %d; want = %d", c.digits, got, want)
		}
	}
}

func BenchmarkAnswer(b *testing.B) {
	answer(3)
}
