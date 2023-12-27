package main

import (
	"fmt"
)

func add(x, y int) int {
	return x + y
}

func multiply (x, y int) int {
	return x * y
}

func aggregateOperation(a, b, c int, arthimetic func(int, int)int) int {
	return arthimetic(arthimetic(a,b),c)
}

func main() {
	firstNo := 5
	secondNo := 4
	thirdNo := 7
	fmt.Println(aggregateOperation(firstNo, secondNo, thirdNo, add))
	fmt.Println(aggregateOperation(firstNo, secondNo, thirdNo, multiply))
}