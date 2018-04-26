package main

import (
	"fmt"
	f "github.com/drafailov/dimitar_rafailov_go_intro/gomath"
)

func main() {
	var a, b int
	fmt.Print("Insert first number: ")
	fmt.Scanln(&a)
	fmt.Print("Insert second number: ")
	fmt.Scanln(&b)
	fmt.Printf("lcm(%d,%d)= ", a, b)
	fmt.Println(f.LCM(a, b))
}
