package main

import "testing"

func TestMySum(t *testing.T) {
	x := mySum(5,6)
	if x != 11 {
		t.Error("Expected ",11," Got ",x)
	}
}
