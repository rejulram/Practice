package main

import (
	"encoding/json"
	"errors"
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
	bs,err := toJson(p)
	if err != nil {
		log.Fatalln("Json Marshal error")
	}
	fmt.Println(string(bs))
}

func toJson(a interface{}) ([]byte,error){
	bs,err := json.Marshal(a)
	if err != nil {
		return []byte{},errors.New(fmt.Sprintln("Error occured",err))
		//Below won't run
//		return []byte{},fmt.Println("Error Occured")
//		Below runs
//      return []byte{},fmt.Errorf("Error Occured",err)
	}
	return bs,nil

}
