package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"


	assertCorrectMesssage(t, repeated, expected)
}

func TestUpperRepeat(t *testing.T) {
	want := UpperrRepeat("a", 10)
	got := "AAAAAAAAAA"

	assertCorrectMesssage(t, want, got)
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i <b.N; i++ {
		Repeat("a", 5)
	}	
}

func ExampleRepeat() {
	repeated := Repeat("a", 5)
	fmt.Println(repeated)
	// Output: aaaaa
}


func assertCorrectMesssage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}