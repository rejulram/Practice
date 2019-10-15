package dog

import (
	"fmt"
	"testing"
)

func TestYears(t *testing.T) {
	ans := Years(10)
	if ans != 70 {
		t.Error("Got:",ans,"Expected : 70")
	}
}

func TestYearsTwo(t *testing.T) {
	ans := YearsTwo(20)
	if ans != 140 {
		t.Error("Got:",ans,"Expected:140")
	}
}

func ExampleYears() {
	fmt.Println(Years(10))
	//Output :
	//70
}

func ExampleYearsTwo() {
	fmt.Println(YearsTwo(20))
	//Output :
	//140
}

func BenchmarkYears(b *testing.B) {
	for i:=0; i < b.N;i++{
		Years(10)
	}
}

func BenchmarkYearsTwo(b *testing.B) {
	for i:=0 ; i <b.N ; i++ {
		YearsTwo(10)
	}
}