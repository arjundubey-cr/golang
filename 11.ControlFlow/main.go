package main

import "fmt"

func main() {
	//if syntax 1
	number := 16
	if number == 16 {
		fmt.Println("Number is 16")
	}

	//if syntax 2
	statePopulations := map[string]int{
		"California": 39250017,
		"Texas":      27862596,
		"Florida":    20612439,
	}
	if pop, ok := statePopulations["Florida"]; ok {
		fmt.Println(pop)
	}
}
