package main

import "fmt"

func main() {
	numberProperties(35)
}

func numberProperties(x int) {
	/* ensures that functionExitLog is called whenever the numberProperties function is exited*/
	defer functionExitLog()

	if (x%2 == 0) {
		fmt.Println("Number is divisible by 2")
		return
	}
	if (x%3 == 0) {
		fmt.Println("Number is divisible by 3")
		return
	}

	if (x%5 == 0) {
		fmt.Println("Number is divisible by 5")
		return
	}

	if (x%7 == 0) {
		fmt.Println("Number is divisible by 7")
		return
	}
}

func functionExitLog() {
	fmt.Println("Function exited")
}