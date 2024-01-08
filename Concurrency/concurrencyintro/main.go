package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	sendEmail("Hello from MK")
}

var email sync.WaitGroup

func sendEmail(message string) {
	
	go func() {
		email.Add(1)
		time.Sleep(time.Millisecond * 250)
		fmt.Printf("Email received: %s\n", message)
		email.Done()
	}()

	fmt.Printf("Email sent: %s\n", message)
	email.Wait()
}