//acdc package contain sum function which add all the values recieved as parameter
package acdc

// Sum function add all the values received as parameter and return the sum
func Sum(xi ...int) int {
	s := 0
	for _, v := range xi {
		s += v
	}
	return s
}
