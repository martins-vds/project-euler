package projecteuler

import (
	"testing"
)

func TestAnswer(t *testing.T) {
	cases := []struct {
		number int
		prime  int
	}{
		{100, 5},
		{13195, 29},
		{600851475143, 6857},
	}

	for _, c := range cases {
		got := answer(c.number)
		want := c.prime

		if got != want {
			t.Errorf("answer(%d) = %d; want = %d", c.number, got, want)
		}
	}
}

func BenchmarkAnswer(b *testing.B) {
	answer(600851475143)
}
