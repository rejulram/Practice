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
	tests := []test{
		test{[]int{10,20,30,40,50},30,},
		test{[]int{1,3,5,7,9,11},6,},
		test{[]int{8,6,3,4,5,10,2,11},6,},
		test{[]int{12,6,5,9,4,2},6},
	}
	for _,v := range tests {
		if r :=CenteredAvg(v.data); r != v.answer {
			t.Error("Got :",r,"Expected",v.answer)
		}
	}
}

func ExampleCenteredAvg() {
	xi := []int{7,2,8,4,1,3}
	fmt.Println(CenteredAvg(xi))
	//Output:
	//4
}

func BenchmarkCenteredAvg(b *testing.B) {
	xi := []int{1,2,3,4,5,6}
		for i:=0 ;i<b.N;i++{
		CenteredAvg(xi)
	}
}
