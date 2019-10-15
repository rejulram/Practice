//Package mymath provides math customized math solutions
package mymath

import "sort"
//CenteredAvg calculate average of list of integers after stripping off smallest and largest element
func CenteredAvg(xi []int) float64 {
	sort.Ints(xi)
	xi = xi[1:len(xi)-1]
	n :=0
	for _,v:= range xi {
		n +=v
	}
	f:= float64(n)/float64(len(xi))
	return f
}
