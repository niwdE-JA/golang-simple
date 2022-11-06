// 03-if-else.go
package main
import "fmt"

func main(){
	var age = 21
	if age>=22 {
		fmt.Println("You're allowed in")
	}else if age>20{
		fmt.Println("You're almost in, come back in a few years")
	}else{
		fmt.Println("You're too young")
	}

}