//Demonstrate variables of several types

package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}

//Variables declared without an explicit initial value are given their zero value. The zero value is:
// `0` for numeric types
// `false` for the boolean type, and
// `""` (the empty string) for strings.

func main() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}

//The expression T(v) converts the value v to the type T. Some numeric conversions:
//var i int = 42
//var f float64 = float64(i)
//var u unit = uint(f)
//Or, put more simply:
//i := 42
//f :=  float64(i)
//u := uint(f)

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}

//When the RHS of the declaration is typed, the new variable is of that same type:
//var i int
//j := i //j is an int

//However, when the RHS contains an untyped numeric constant, the new variable may be an int, float64, or complex128 depdending
//on the precision of the constant:

//i := 42 //int
//f := 3.142 //float64
//g := 0.867 + 0.5i //complex128

func main() {
	v := 42 // change me!
	fmt.Printf("v is of type %T\n", v)
}

//Constants are declared like variables, but with the const keyword
//Constants can be character, string, boolean, or numeric values
//Constants cannot be declared using the := syntax

const Pi = 3.14

func main() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
