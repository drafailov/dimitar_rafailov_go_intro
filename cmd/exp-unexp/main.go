package main

import (
	"fmt"
	"github.com/drafailov/dimitar_rafailov_go_intro/govisibility"
)

func main() {
	//Print a constant of the exported type
	fmt.Println(govisibility.Con)
	// con is not visible
	//fmt.Print(govisibility.con)
	// Newnum is visible
	number := govisibility.Newnum(56)
	fmt.Println(number)
	sqr := govisibility.Shape{
		Name: "Square",
		Area: 30,
		//perimeter is not visible
		//perimeter:15,
	}
	fmt.Println(sqr)
	// Creating an object Square with fields Name, Aria & perimetar
	// are visible from unexported struct shape
	sqr1 := govisibility.Square{
		Perimeter: 15,
	}
	sqr1.Name = "Square"
	sqr1.Area = 33
	fmt.Print(sqr1)
}
