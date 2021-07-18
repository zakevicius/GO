package hello

import "testing"

func TestHello(t *testing.T) {
	want := "Hi, Somebody. Welcome!"
	if got := Hello("Somebody"); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}
