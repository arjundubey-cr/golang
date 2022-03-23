package main

import "fmt"

func main() {
	a := "start"
	defer fmt.Println(a)
	a = "end"

	fmt.Println("start")
	//panic("something bad happened")
	fmt.Println("end")
}
