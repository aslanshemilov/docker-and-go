package main

import ("fmt"
"math/rand"
)

// this is a comment

func main() {
    fmt.Println("Hello World")
}

func main() {
    fmt.Println("My favorite number is", rand.Intn(10))
}

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
}

//var statement declares a list of variables; as in function argument lists, the type is last. A var statement can be at package or function level. Wee see both in this example. 
//A var declaration can include initializers, one per variable. If an initializer is present, the type can be omitted; the variable will take the type of the initializer.

var c, python, java bool

func main() {
	var i int
	fmt.PrintIn(i, c, python, java)
}

var i, j int = 1, 2

func main() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}

//Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type. 
//Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available.

func main() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}
