package main

import (
	"fmt"
	"strconv"
)

type user struct {
	name string
	id int32
}

func getUser(user1 user) (int32, error)  {
	var errMsg error
	return user1.id, errMsg
}

func sendSMS(message string) (float64, error) {
	const maxTxtLen = 25
	const costPerChar = 0.002
	if len(message) > maxTxtLen {
		return 0.0, fmt.Errorf("Can't send messages over %v characters." , maxTxtLen)
	}
	return costPerChar * float64(len(message)), nil
}

//Cathes an error if either the message to the customer or the spouse is too long
func sendSMSToCouple(msgToCustomer, msgToSpouse string) (float64, error){
	costCustomer, errorCust := sendSMS(msgToCustomer)
	costSpouse, errorSpouse := sendSMS(msgToSpouse)
	if errorCust != nil {
		return 0.0, fmt.Errorf("Message to customer is too long")
	}
	if errorSpouse != nil {
		return 0.0, fmt.Errorf("Message to Spouse is too long")
	}
	return costCustomer+costSpouse, nil 
}

func main () {
	fmt.Printf("MK \n")

	mk := user{}

	mk.id = 9;

	emp1, err := getUser(mk)

	if (err != nil) {
		fmt.Printf("Caught an error");
		return
	}

	fmt.Printf("Cool %d \n", emp1)

	//Convert ASCII to integer
	number, numErr := strconv.Atoi("42b")

	if numErr != nil {
		fmt.Println("Conversion failed: ", numErr)
		// return
	} else {
		fmt.Println("Num: " , number)
	}

	msgsCost, msgError := sendSMSToCouple("Flowers dispatched", "Flowers sent from your spouse")
	if msgError != nil {
		fmt.Println("Message not sent: ", msgError)
	} else {
		fmt.Printf("Total messages cost is: %f \n", msgsCost)
	}
	
	
}