package projecteuler

import (
	"testing"
)

func TestAnswer(t *testing.T) {
	cases := []struct {
		limit          int
		smallestNumber int
	}{
		{10, 2520},
		{20, 232792560},
	}

	for _, c := range cases {
		got := answer(c.limit)
		want := c.smallestNumber

		if got != want {
			t.Errorf("answer(%d) = %d; want = %d", c.limit, got, want)
		}
	}
}

func BenchmarkAnswer(b *testing.B) {
	answer(20)
}
