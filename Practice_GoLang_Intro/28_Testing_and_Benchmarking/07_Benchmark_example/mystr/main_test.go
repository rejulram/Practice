package mystr

import (
	"fmt"
	"strings"
	"testing"
)

func TestCat(t *testing.T) {
	s := "You are doing great"
	xs := strings.Split(s, " ")
	s = Cat(xs)
	if s != "You are doing great" {
		t.Error("Got:", s, "Expected", "You are doing great")
	}
}

func TestJoin(t *testing.T) {
	s := "You are doing great"
	xs := strings.Split(s, " ")
	s = Join(xs)
	if s != "You are doing great" {
		t.Error("Got:", s, "Expected", "You are doing great")
	}
}

func ExampleCat() {
	s := "You are doing great"
	xs := strings.Split(s, " ")
	fmt.Println(Cat(xs))
	//Output:
	//You are doing great
}

func ExampleJoin() {
	s := "You are doing great"
	xs := strings.Split(s, " ")
	fmt.Println(Join(xs))
	//Output:
	//You are doing great
}

const s = "There are many languages spoken in Kerala, though the official language of Kerala is Malayalam. This language belongs to the Dravidian group of languages. Almost 90% people of Kerala speak Malayalam. There are 5 main regional dialects of Malayalam. Other common languages that are spoken in Kerala are English and Tamil. These Kerala languages are often interspersed with words from Sanskrit, Latin, Urdu, etc. due to foreign influence"

var xs []string

func BenchmarkCat(b *testing.B) {
	xs = strings.Split(s, " ")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Cat(xs)
	}
}

func BenchmarkJoin(b *testing.B) {
	xs = strings.Split(s, " ")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Join(xs)
	}
}
