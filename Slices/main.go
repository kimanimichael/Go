package main

import (
	"errors"
	"fmt"
)

const (
	planFree = "free"
	planPro = "pro"
)

func getMessageWithRetries () [3]string {
	var messageReplies [3]string
	messageReplies[0] = "\nClick here to sign up\n"
	messageReplies[1] = "Pretty please click here\n"
	messageReplies[2] = "We demand that you sign up\n"
	return messageReplies
}

func getMessagesForPlan(plan string) ([]string, error) {
	allMessages := getMessageWithRetries()
	if plan == planPro {
		return allMessages[:], nil
	} else if plan == planFree {
		return allMessages[0:2], nil
	} else {
		return nil, errors.New("Unsupported plan")
	}
}

func main()  {
	fmt.Printf("MK\n")
	fmt.Println(getMessageWithRetries()[1])

	proMessages,_ := getMessagesForPlan("pro")
	fmt.Println(proMessages[:])

	freeMessages,_ := getMessagesForPlan("free")
	fmt.Println(freeMessages[:])
}