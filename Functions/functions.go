package main

import (
	"errors"
	"fmt"
)

// Type comes after variable in the function signature
func concat(temp string, user string) string {
	return temp + user
}

// With multiple variables the type can be listed after the last one
func add(x, y int) int {
	return x + y
}

func main() {
	holder := "Message from "
	texter := "MK"
	fmt.Println(concat(holder, texter))

	sendsSoFar := 10
	sendsToAdd := 2
	sendsSoFar = incrementSends(sendsSoFar, sendsToAdd)

	fmt.Printf("Total sends = %v \n", sendsSoFar)

	name1, name2 := getNames()

	fmt.Printf(name1 + " " + name2 +"\n")

	duration := 2
	kms := 150
	//Ignore one return value
	ukspeed, _, calcError := getSpeed(duration, kms)

	fmt.Printf("UK speed: %v\n", ukspeed)
	fmt.Printf("Error: %v\n", calcError)
}

func incrementSends(currentSends, newSends int) int {
	currentSends = currentSends + newSends
	return currentSends
}

func getNames() (string, string) {
	return "James", "Jude"
}

func getSpeed(time, distance int) (float64, float64, error) {
	var kmph float64 
	var errorMsg error
	kmph, errorMsg = divide(distance, time)
	kmtom := 0.621371
	mph := kmph * kmtom
	return kmph, mph, errorMsg
}

//Implement early returns
func divide(dividend, divisor int)(float64, error) {
	if divisor == 0 {
		return 0, errors.New("Division by zero is undefined")
	}

	var quotient float64
	quotient = float64(dividend)/float64(divisor)
	return quotient, nil
}