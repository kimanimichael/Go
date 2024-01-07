package main

import (
	"fmt"
	"strings"
)

func main() {
	testMessage := "Fuck this shit. I am better"
	
	fmt.Println(testMessage)
	removeProfanity(&testMessage)
	/* should cause a crash unless we check for a nil pointer in the removeProfanity function */
	removeProfanity(nil)
	fmt.Println(testMessage)

}
/* takes in a pointer as a parameter*/
func removeProfanity(message *string) {
	if message == nil {
		fmt.Println("Pointer is nil")
		return
	}
	/* de-reference the pointer and store the value in messageVal */
	messageVal := *message
	messageVal = strings.ReplaceAll(messageVal, "Fuck", "****")
	messageVal = strings.ReplaceAll(messageVal, "shit", "****")
	/* point the original pointer to the newly mutated variable */
	*message = messageVal
}