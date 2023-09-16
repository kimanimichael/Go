package main

import "fmt"

func main() {
	fmt.Println("Welcome to Go")

	var username string = "mk"
	var passworkd string  = "5678765"

	fmt.Println("Authorization: Basic" , username +":"+passworkd)
}