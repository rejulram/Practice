// Package dog allows us to fully understand dogs
package dog

//Years convert human years to dog years .
func Years(n int) int {
	return n*7
}

//YearsTwo convert human years to dog years .
func YearsTwo(n int) int {
	count :=0
	for i:=0;i< n ; i++ {
		count += 7
	}
	return count
}