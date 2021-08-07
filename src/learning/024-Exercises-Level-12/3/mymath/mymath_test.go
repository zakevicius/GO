package mymath

import (
	"fmt"
	"testing"
)

func TestCenteredAvg(t *testing.T) {
	type test struct {
		data []int
		answer float64
	}
	
	tests := []test {
		{[]int{1, 2, 4, 5}, 3},
		{[]int{1, 2, 5, 3}, 2.5},
		{[]int{10, 2, 4, 5}, 4.5},
		{[]int{10, 20, 4, 6}, 8},
	}

	for _, v := range tests {
		res := CenteredAvg(v.data)
		if res != v.answer {
			t.Error("Expected", v.answer, ", got", res)
		}
	}
}

func ExampleCenteredAvg() {
	fmt.Println(CenteredAvg([]int{1, 2, 4, 5}))
	// Output: 3
}

func BenchmarkCenteredAvg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CenteredAvg([]int{1, 2, 4, 5})
	}
}
