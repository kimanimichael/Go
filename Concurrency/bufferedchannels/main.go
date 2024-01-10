package main

import "fmt"

func addEmailsToQueue(emails []string) chan string {
	emailsToSend := make(chan string, len(emails))
	// go func ()  {
		
	// }()
	for _, email := range emails {
		emailsToSend <- email
	}
	return emailsToSend
}

func sendEmails(batchSize int, ch chan string) {
	for i := 0; i < batchSize; i++ {
		email := <-ch
		fmt.Printf("Sending email: %s\n", email)
	}
}

func test(emails ...string) {
	fmt.Printf("Adding %d emails to queue...\n", len(emails))
	ch := addEmailsToQueue(emails)
	fmt.Printf("Sending emails...\n")
	sendEmails(len(emails), ch)
	fmt.Println("==============================")
}

func main() {
	test("Hello MK", "Welcome MK", "Be nice MK")
}