package projecteuler

import (
	"testing"
)

func TestAnswer(t *testing.T) {
	cases := []struct {
		nthTerm     int
		primeNumber int
	}{
		{6, 13},
		{10001, 104743},
	}

	for _, c := range cases {
		got := answer(c.nthTerm)
		want := c.primeNumber

		if got != want {
			t.Errorf("answer(%d) = %d; want = %d", c.nthTerm, got, want)
		}
	}
}

func BenchmarkAnswer(b *testing.B) {
	answer(10)
}
