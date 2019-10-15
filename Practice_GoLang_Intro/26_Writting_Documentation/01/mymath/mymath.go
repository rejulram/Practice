// Package mymath provides ACME math solution
package mymath

// Sum adds an unlimited values of type int
func Sum(xi ...int) int {
	sum :=0
	for _,v := range xi {
		sum += v
	}
	return sum
}