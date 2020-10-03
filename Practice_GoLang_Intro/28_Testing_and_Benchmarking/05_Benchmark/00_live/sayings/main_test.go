package sayings

import (
	"fmt"
	"testing"
)

func TestGreet(t *testing.T) {
	s := Greet("Hi")
	if s != "Hello Rejul" {
		t.Error("Got :", s, "Expected :", "Hello Rejul")
	}
}

func ExampleGreet() {
	fmt.Println(Greet("Hi"))
	//Output:
	//Hello Rejul
}

func BenchmarkGreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Greet("Hi")
	}
}
