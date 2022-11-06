// 07-structs.go
package main

import "fmt"

func main(){
	type Person struct{
		name string
		age int
		gender bool
	}

	var john = Person{
		name: "John Doe",
		age: 21,
		gender : false,
	}

	fmt.Println("My 'Person' struct :\n", john)

	john.age++ // we can change values in struct fields
	fmt.Println("My 'Person' struct :\n", john.name, john.age, john.gender)

	// other ways to initilize structs...
	guy := Person{"Guy Moon", 36, true} // here, the fields must be entered in the same order they were defined, since their names are not specified

	jane := Person{} // here, jane would be initialized with default values
	fmt.Println(guy, jane)


	//define struct 'anonymously' w/o type name
	zino := struct{
		name string
		label string
		headiesWon int
	} { "Zinoleesky", "Marlian records", 0}

	fmt.Println("Anonymous struct :", zino)
}