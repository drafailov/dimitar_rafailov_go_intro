package main

import (
	"fmt"
	gcd "./divisor"
)

func main() {
	var a, b int
	fmt.Print("Insert first number: ")
	fmt.Scanln(&a)
	fmt.Print("Insert second number: ")
	fmt.Scanln(&b)
	fmt.Printf("gcd(%d,%d)= ", a, b)
	fmt.Println(gcd.GCD(a, b))
}
