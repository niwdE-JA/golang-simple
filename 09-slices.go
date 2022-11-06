// 09-slices.go
package main

import "fmt"

func main(){
	arr := [5]float32{1,2,1,3,2}
	
	mySlice := arr[1:4] // A slice is like an array with special functions like append()
	mySlice = append(mySlice, 888)
	fmt.Println(mySlice)

	//other way of defining slice
	arr2 := []float32{1,2,3}

	//append two slices together
	mySlice2 := append(mySlice, arr2...) // '...' is a spread operator
	fmt.Println(mySlice2)
}