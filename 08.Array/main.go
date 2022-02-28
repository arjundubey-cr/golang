package main

import "fmt"

func main() {
	arr := [...]int{1, 2, 3}
	fmt.Println(arr)

	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := a[:]  //slice all of the elements
	c := a[3:] //slice from 4th to the end
	d := a[:6] //slice first 6
	e := a[3:6]
	f := append(a[:2], a[3:]...)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)

	fmt.Println(f)
	fmt.Println(a) //unexpected behaviour here
}
