//Package govisibility contains functions, variables, constants & etc.
//showing difference between exported and unexported
package govisibility

// Con is a constant from exported type
const Con = 5

// con is a constant from unexported type
const con = 8

// Num is from exported type
type Num int

//num is unxported type
type num int

//Newnum is exported function begins with capitalized letter
//returns unexported type num
func Newnum(value int) num {
	return num(value)
}

//newnum is unexported function begins with small letter
func newnum(value int) num {
	return num(value)
}

//Shape is struct with exported fields (Name & Aria) & unexported perimeter
type Shape struct {
	Name      string
	Area      int
	perimeter int
}

// shape is unexported struct
type shape struct {
	Name string
	Area int
}

//Square is exported struct
type Square struct {
	shape
	Perimeter int
}
