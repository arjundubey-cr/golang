package main

import "fmt"

func main() {
	statePopulation := map[string]int{
		"California": 39250017,
		"Texas":      27862596,
		"Florida":    20612439,
	}
	fmt.Println(statePopulation)

	statePopulation["Georgia"] = 10310371
	fmt.Println(statePopulation)

	delete(statePopulation, "Georgia")

	_, ok := statePopulation["Oho"]
	fmt.Println(ok)

	sp := statePopulation
	delete(sp, "Florida")
	fmt.Println(statePopulation)
}
