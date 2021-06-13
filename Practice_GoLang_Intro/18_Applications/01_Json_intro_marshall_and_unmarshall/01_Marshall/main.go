package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main(){
	type ColourGroup struct {
		Id int
		Name string
		Colors []string
	}
	group := ColourGroup{
		Id: 1,
		Name: "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}
	b,err := json.Marshal(group)
	if err != nil{
		fmt.Println("error",err)
	}
	os.Stdout.Write(b)
}
