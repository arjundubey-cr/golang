package main

import (
	"fmt"
)

type Doctor struct {
	number     int
	actorName  string
	companions []string
}

type Animal struct {
	Name   string
	Origin string
}
type Bird struct {
	Animal
	SpeedKPH float32 `required:"true" max:"100"`
	CanFly   bool
}

func main() {
	aDoctor := Doctor{
		number:    3,
		actorName: "Jon Pertwee",
		companions: []string{
			"Liz Shaw",
			"Jo Grant",
			"Sarah Jane Smith",
		},
	}
	fmt.Println(aDoctor)

	//Inline structs
	aNewDoctor := struct{ name string }{name: "John Doe"}
	fmt.Println(aNewDoctor)

	//Composition relationship to replicate inheritance
	b := Bird{}
	b.Name = "Emu"
	b.Origin = "Australia"
	b.SpeedKPH = 48
	fmt.Println(b.Name)
}
