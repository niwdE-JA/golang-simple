// 06-functions.go
package main
import "fmt"


func countDown(num int){
	for i := num; i>0; i-- {
		fmt.Println(i, "units left")
	}
	fmt.Println("Countdown finished.")
}

func add(num1 int, num2 int) int{
	var sum = num1 + num2
	return sum
}

func main(){
	countDown(8)
	countDown(2)
	// countDown(3.33) //--throws an error

	fmt.Println("Sum of 2 and 3 :", add(2,3))

	newName := add

	fmt.Println("Sum of 2 and 3 :", newName(2,3))

}