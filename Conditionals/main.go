package main

import "fmt"

func main() {
	// Basic if else
	messageLen := 6
	maxMessageLen := 10
	if messageLen < maxMessageLen {
		fmt.Println("Message sent")
	} else if messageLen == maxMessageLen {
		fmt.Println("Message sent but there's a risk of overflow")
	} else {
		fmt.Println("Message not sent. Overflow")
	}

	// If condition with variable living only inside an if scope
	// emailLength only lives within this scope
	if emailLength := 10; emailLength <= maxMessageLen {
		fmt.Println("Email sent")
	}
}