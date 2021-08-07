package main

import "testing"

func TestMySum(t *testing.T) {
	type test struct {
		data []int
		answer int
	}

	tests := []test{
		{[]int{1, 2, 3, 4}, 10},
		{[]int{3, 4}, 7},
		{[]int{2, 14}, 16},
	}

	for _, test := range tests {
		x := mySum(test.data...)

		if x != test.answer {
			t.Error("Expected", test.answer, ", got", x)
		}
	}
}
