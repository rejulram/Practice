package word

import (
	"UdemyClass/Practice/29_Ninja_Level_13/02/quote"
	"fmt"
	"testing"
)

func TestUseCount(t *testing.T) {
	m := UseCount("one two three three three")
	for k, v := range m {
		switch k {
		case "one":
			if v != 1 {
				t.Error("Got:", v, "Expected 1")
			}
		case "two":
			if v != 1 {
				t.Error("Got:", v, "Expected 1")
			}
		case "three":
			if v != 3 {
				t.Error("Got:", v, "Expected 3")
			}
		}
	}
}

func TestCount(t *testing.T) {
	n := Count("one two three")
	if n != 3 {
		t.Error("Got:", n, "Expected : 3")
	}
}

func ExampleCount() {
	fmt.Println(Count("one two three"))
	//Output :
	//3
}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count(quote.SunAlso)
	}

}

func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount(quote.SunAlso)
	}
}
