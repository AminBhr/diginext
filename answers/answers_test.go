package answers

import (
	"reflect"
	"testing"
)

type inputs struct {
	in  string
	l   int
	out string
	n   int
}

type swapInputs struct {
	in  []int
	out []int
	n   int
}

func TestRepeatedString(t *testing.T) {
	ts := []inputs{
		{in: "abcac", l: 10, out: "abcacabcac", n: 4},
		{in: "ddb", l: 5, out: "ddbdd", n: 0},
		{in: "ghca", l: 9, out: "ghcaghcag", n: 2},
	}

	for _, tc := range ts {
		r, c := RepeatedString(tc.in, tc.l)
		if r != tc.out || c != tc.n {
			t.Error()
			continue
		}
	}
}

func TestAlternatingString(t *testing.T) {
	ts := []inputs{
		{in: "AABAAB", out: "ABAB", n: 2},
		{in: "AAAA", out: "A", n: 3},
		{in: "BBBBB", out: "B", n: 4},
		{in: "BABABA", out: "BABABA", n: 0},
		{in: "ABABABAB", out: "ABABABAB", n: 0},
		{in: "AAABBB", out: "AB", n: 4},
	}

	for _, tc := range ts {
		r, c := AlternatingString(tc.in)
		if r != tc.out || c != tc.n {
			t.Error()
			continue
		}
	}
}

func TestSwapSorting(t *testing.T) {
	ts := []swapInputs{
		{in: []int{7, 1, 3, 2, 4, 5, 6}, out: []int{1, 2, 3, 4, 5, 6, 7}, n: 5},
		{in: []int{1, 2, 4, 3, 5}, out: []int{1, 2, 3, 4, 5}, n: 1},
		{in: []int{4, 3, 2, 1}, out: []int{1, 2, 3, 4}, n: 2},
		{in: []int{3, 5, 4, 2, 1}, out: []int{1, 2, 3, 4, 5}, n: 4},
	}

	for _, tc := range ts {
		r, c := SwapSorting(tc.in)
		if !reflect.DeepEqual(r, tc.out) || c != tc.n {
			t.Error()
			continue
		}
	}
}
