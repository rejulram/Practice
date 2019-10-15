package main

import (
	"UdemyClass/Practice/29_Ninja_Level_13/02/quote"
	"UdemyClass/Practice/29_Ninja_Level_13/02/word"
	"fmt"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))
	for k,v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v,k)
	}
}
