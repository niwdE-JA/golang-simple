package main
import "fmt"

func main(){
	animal :="dog"
	switch animal{
		case "cat":
			fmt.Println("Your animal is a cat!")
		case "eagle":
			fmt.Println("You have selected an eagle!")
		default:
			fmt.Println("I don't recognize this animal")
	}
}