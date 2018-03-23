//Go has only one looping construct, the for loop.
//The basic for loop has three componenets separated by semicolons:

//init statement: executed before the first iteration; short variable declaration, variables declared there are visible only in the scope of the for statement
//condition expression: evaluated before every iteration
//post statement: executed at the end of every iteration

//The loop will stop iterating once the boolean condition evaluates to false.

package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ { //init statement; condition expression; post statement
		sum += i
	}
	fmt.Println(sum)
}

//The init and post statement are optional.

func main() {
	sum := 1 //initialization
	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)
}

//At that point you can drop the semicolons: C's while is spelled for in Go

func main() {
	sum := 1 //initialization
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

//If you omit the loop condition it loops forever, so an infinite loop is compactly expressed.

func main() {
	for {
	}
} //received output "process took too long"


//Go's if statements are like its for loops; the expression need not be surrounded by parentheses () but the braces {} are required.

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
}
