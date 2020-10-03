package main

import (
	"UdemyClass/Practice/28_Testing_and_Benchmarking/07_Benchmark_example/mystr"
	"fmt"
	"strings"
)

const s = "There are many languages spoken in Kerala, though the official language of Kerala is Malayalam. This language belongs to the Dravidian group of languages. Almost 90% people of Kerala speak Malayalam. There are 5 main regional dialects of Malayalam. Other common languages that are spoken in Kerala are English and Tamil. These Kerala languages are often interspersed with words from Sanskrit, Latin, Urdu, etc. due to foreign influence"

func main() {
	xs := strings.Split(s, " ")
	for _, v := range xs {
		fmt.Println(v)
	}
	fmt.Printf("\n%s\n", mystr.Cat(xs))
	fmt.Printf("\n%s\n\n", mystr.Join(xs))
}
