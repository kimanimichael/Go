package main

import (
	"fmt"
)

func printPrimeNumbers(upperLimit int) {
	for n:=2; n < upperLimit + 1;n++ {
		if n == 2 {
			fmt.Println(n)
			continue
		}
		if n%2 == 0 {
			continue
		}
		isPrime := true
		for i:=3; i*i < (n+1); i++ {
			if n%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			fmt.Println(n)
		}
	}
}

func main() {
	fmt.Printf("MK\n")
	printPrimeNumbers(10)
}