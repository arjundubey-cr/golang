package main

import "fmt"

func main() {
	// keyword variable_name type
	var username string = "hitesh"
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255 //uint8>=256 is out of bounds for uint8;
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float32 = 255.45445112554454546 //uint8>=256 is out of bounds for uint8;
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	var largeFloat float64 = 255.45445112554454546 //uint8>=256 is out of bounds for uint8;
	fmt.Println(largeFloat)
	fmt.Printf("Variable is of type: %T \n", largeFloat)

	var anotherVariable int //uint8>=256 is out of bounds for uint8;
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	var tempVariable = "this variable"
	fmt.Println(tempVariable)
	fmt.Printf("Variable is of type: %T \n", tempVariable)

	newVariable := "newVar"
	fmt.Printf("Variable is of type: %T \n", newVariable)

}
