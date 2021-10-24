package projecteuler

import (
	"testing"
)

func TestAnswer(t *testing.T) {
	cases := []struct {
		nthTerm int
		sum     int
	}{
		{4000000, 4613732},
	}

	for _, c := range cases {
		got := answer(c.nthTerm)
		want := c.sum

		if got != want {
			t.Errorf("answer(%d) = %d; want = %d", c.nthTerm, got, want)
		}
	}
}

func BenchmarkAnswer(b *testing.B) {
	answer(4000000)
}
