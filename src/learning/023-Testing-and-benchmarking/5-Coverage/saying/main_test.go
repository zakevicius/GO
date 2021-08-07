package saying

import (
	"fmt"
	"testing"
)

func TestGreet(t *testing.T) {
	s := Greet("Foo")
	if s != "Welcome Foo" {
		t.Error("Expected", "Welcome Foo", "got", s)
	}
}

func ExampleGreet() {
	fmt.Println(Greet("Foo"))
	// Output
	// Welcome Foo
}

func BenchmarkGreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Greet("James")
	}
}
