// 05-loops.go
package main
import "fmt"

func main() {
	// for loop
	for i:=0; i<10; i++ {
		fmt.Println("For loop :", i)
	}

	// while loop
	j:=0
	for j<10{
		fmt.Println("While loop :", j)
		j++
	} 

	// using break with no condition in for
	k:=0
	for {
		if k>=10 {
			break
		}
		fmt.Println("for loop with break :", k)
		k++
	}

	// 'continue' and 'break' keywords are defined in the usual way 
}