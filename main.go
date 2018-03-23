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

