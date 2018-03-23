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
