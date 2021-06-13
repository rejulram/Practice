package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main(){
	password1 := `password123`
	bs,err := bcrypt.GenerateFromPassword([]byte(password1),bcrypt.MinCost)
	if err != nil{
		fmt.Println(err)
	}
	password2 := `password123`
	err = bcrypt.CompareHashAndPassword(bs,[]byte(password2))
	if err != nil{
		fmt.Println("Password 2 You can't Login")
		return
	}
	fmt.Println("Password 2 Logged in")

	password3 := `password124`
	err = bcrypt.CompareHashAndPassword(bs,[]byte(password3))
	if err != nil{
		fmt.Println("Password3 You can't Login")
		return
	}
	fmt.Println("Password 3 Logged in")	

}
