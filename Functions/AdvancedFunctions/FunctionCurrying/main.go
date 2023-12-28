package main

import "fmt"

func main() {
	squareFunction := selfMath(multiply)
	doubleFunction := selfMath(add)

	fmt.Println(squareFunction(16))
	fmt.Println(doubleFunction(19))
}

func add(x, y int) int {
	return x + y
}

func multiply(x, y int) int {
	return x * y
}
/* currying function that returns a new function */
func selfMath(mathFunc func(int, int)int)func(int)int {
	return func(x int) int{
		return mathFunc(x, x)
	}
}