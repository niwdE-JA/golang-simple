// 08-arrays.go
package main

import "fmt"

func main(){
	// array elements must allhave the same type
	purchases := [4]float32{12.3, 3.9,1.5,4.87}
	fmt.Println(purchases) 

	//declaration w/o assigning
	var sales [4]float32
	sales[1] = 22.21 
	fmt.Println(sales) 
	
	//alternative array declaration
	heights := [...]float32{2,3,4,1,2,1} //this '[...]' syntax does NOT work with var 

	len(heights) //can be used in loops (NOT heights.length)
}