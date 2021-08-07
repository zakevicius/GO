package word

import (
	"fmt"
	"testing"
)

const s = "One one two three two"

func TestCount(t *testing.T) {
	total := Count(s)

	if total != 5 {
		t.Error("Expected 5, got", total)
	}
}

func TestUseCount(t *testing.T) {
	res := UseCount(s)

	if res["One"] != 1 || res["one"] != 1 || res["two"] != 2 || res["three"] != 1 {
		t.Errorf("Expected One = 1 one = 1 two = 2 three = 1, got One = %d one = %d two = %d three = %d", res["One"], res["one"], res["two"], res["three"])
	}
}

func ExampleCount() {
	fmt.Println(Count("One one two three two"))
	// Output: 5
}

func ExampleUseCount() {
	fmt.Println(UseCount("One one two three two"))
	// Output: map[One:1 one:1 three:1 two:2]
}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count(s)
	}
}

func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount(s)
	}
}
