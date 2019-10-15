package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct{
	First string
	Last string
	Sayings []string
}

func main() {
	p := person{
		First:"Rejul",
		Last:"Ramakrishnan",
		Sayings:[]string{"Hi","Hello","How are you",},
	}
	bs,err := json.Marshal(p)
	if err != nil {
		log.Fatalln("Json Marshal error")
	}
	fmt.Println(string(bs))
}
