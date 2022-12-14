// 01-variables.go
package main

import "fmt"

func main(){
	var name = "John Doe"
	fmt.Println("Hello, World!", name)
	name = "Jack Sparrow"
	fmt.Println("Hello, World!", name)
	// name = 12 //raises error because of type
	
	var age int = 70 //specifying the type explicitly(not so useful here where it's obvious)
	fmt.Println("age is :", age)

	var pet string //specify type explicitly w/o assignment
	pet = "Dog"
	fmt.Println("Pet is :", pet)

	//variables, when declared w/o initialization, are assigned defaul values(0, "", false, etc)

	//compound creation :
	var i, j, k = 3, 2, 8 

	//block creation
	var (
		a = 3
		b = 8
		c = 1
	)

	fmt.Println(i,j,k,a,b,c)

	school := "unilag" //alternative to var
	 
	//constant value
	const unchanging = "This value is immutable"
	
	// This would not work: 
	// unchanging = "something else"
}