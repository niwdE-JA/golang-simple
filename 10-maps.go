// 10-maps.go
package main

import "fmt"

func main(){
	// maps can be created using make()
	cart := make([string] int) //[key] value

	cart["book"] = 2
	cart["laptop"] = 1
	cart["mouse"] += 1 // assigns default value of 0, then increments to 1
	fmt.Println(cart)
	
	//READ from map
	value1 := cart["book"] // if not found, value1 sets to default
	value2, optionalIsFoundBool = cart["invalidKey"] // optionalIsFoundBool tells us if value with specified key exists
	fmt.Println(value1)
	fmt.Println(value2, optionalIsFoundBool)

	// delete key-value pairs from map
	delete(cart, "mouse")

	// NOTE: maps are NOT ordered, so iterating over them may produce an unpredictable order
	

}