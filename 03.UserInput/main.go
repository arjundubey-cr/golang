package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcomeMessage := "Hello, welcome to the application."
	fmt.Println(welcomeMessage)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our pizza")

	//comma ok || err ok syntax
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating, ", input)
	fmt.Printf("Type of this rating is %T \n", input)
}
