package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleLetter() {
	Seed(11)
	fmt.Println(Letter())
	// Output: g
}

func BenchmarkLetter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Letter()
	}
}

func ExampleLexify() {
	Seed(11)
	fmt.Println(Lexify("?????"))
	// Output: gbrma
}

func BenchmarkLexify(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Lexify("??????")
	}
}

func ExampleSlice() {
	Seed(11)
	fmt.Println(Slice([]string{"happy", "times", "for", "everyone", "have", "a", "good", "day"}))
	// Output: happy
}

func BenchmarkSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Slice([]string{"happy", "times", "for", "everyone", "have", "a", "good", "day"})
	}
}