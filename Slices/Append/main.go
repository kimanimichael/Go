package main

import (
	"fmt"
)

func createMessageSlice () []string {
	var messageStrings [3]string
	messageStrings[0] = "Click here to sign up"
	messageStrings[1] = "Pretty please click here"
	messageStrings[2] = "We demand that you sign up"

	messageSlice := messageStrings[:]
	return messageSlice
}

func main() {
	messageSlice := createMessageSlice()
	fmt.Println(messageSlice)
	/* test append function */
	messageSlice = append(messageSlice, "MK")
	fmt.Println(messageSlice)
}