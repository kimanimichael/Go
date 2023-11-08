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
	messageReplies[0] = "Click here to sign up"
	messageReplies[1] = "Pretty please click here"
	messageReplies[2] = "We demand that you sign up"
	return messageReplies
}

/* function returning a slice */
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

	/* Print array elements with quotes */
	fmt.Printf("%+q", proMessages[:])

	fmt.Printf("\n")

	freeMessages,_ := getMessagesForPlan("free")
	fmt.Printf("%v \n", freeMessages[:])

	/* print each element of a slice with a for loop */
	for i:=0; i<len(freeMessages); i++ {
		fmt.Println(freeMessages[i])
	}

	/* create a slice of size 5 with an underlying array of size 10 */
	/* growing a slice past the size of the underlying array leads to reallocation of the underlying array*/
	definedArraySizeSlice := make([]int, 5, 10)
	fmt.Println(definedArraySizeSlice[:])
	

	/* create a slice of size 5 with an underlying array of size 5 */
	implicitArraySizeSlce := make([]int, 5)
	fmt.Println(implicitArraySizeSlce[:])

	/* functions for checking length and capacity*/
	fmt.Println(cap(definedArraySizeSlice))
	fmt.Println(len(implicitArraySizeSlce))

	messagesCost := getMessagesCost(proMessages)
	fmt.Println(messagesCost)



}

func getMessagesCost(messages []string) []float64 {
	messagesCost := make([]float64, len(messages))
	for i:=0; i < len(messages); i++ {
		message := messages[i]
		cost := float64(len(message)) * 0.01
		messagesCost[i] = cost
	}
	return messagesCost
}

}