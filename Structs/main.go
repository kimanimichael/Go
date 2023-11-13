package main

import "fmt"

type car struct {
	make string
	brand string
	wheels int
	power_sources int
	// Doors an anonymous struct within the struct car
	doors struct {
		number int
		material string
	}
}

type user struct {
	name string
	number int
}

type messageToSend struct {
	message string
	// Nested structs
	sender user
	recepient user
}

type credentialsToStore struct {
	// Embedded struct - Fields in user become top level fields of credentialsToStore
	user
	password string
	location string
}

func validate(message messageToSend) bool{
	if message.recepient.name == "" {
		return false
	}
	if message.recepient.number == 0 {
		return false
	}
	if message.sender.name == "" {
		return false
	}
	if message.sender.number == 0 {
		return false
	}
	return true
}

func sendMessage(payload messageToSend) string{
	if validate(payload) == false {
		outcome := "Validation failed"
		fmt.Println(outcome)
		return outcome
	}
	fmt.Printf("Message sent to %v \n", payload.recepient.name)
	outcome := "Success"
	return outcome
}

type sales_data struct{
	initial_stock float64
	added_stock float64
	sold_stock float64
	remaining_stock float64
}

// Method in Go mainly - Example in a struct
func (data sales_data) stockDiscrepancy() float64{
	discrepancy := data.initial_stock + data.added_stock - data.sold_stock - data.remaining_stock
	return discrepancy
}

type authenticationInfo struct {
	username string
	password string
}

// Method in Go - Example 2
func (auth authenticationInfo) getBasicAuth() string {
	var basicAuth string
	basicAuth = "Authorization: Basic " + auth.username + ":" + auth.password
	return basicAuth
}


func main()  {
	// Creates a new empty car called myCar - All members initialized to their default values
	myCar := car{}
	myCar.power_sources = 2
	myCar.doors.number = 4
	fmt.Printf("Number of doors in my car: %v \n", myCar.doors.number)

	dadsCar := sales_data{
		100.9,
		150.7,
		300.7,
		3.0}

	fmt.Println(dadsCar.remaining_stock)

	newMessage := messageToSend{}
	newMessage.sender.name = "MK"
	newMessage.recepient.name = "Brexit"
	newMessage.sender.number = 7
	newMessage.recepient.number = 1

	sendMessage(newMessage)

	// Anonymous struct - Can also be created within a struct (Checkout nested structs)
	myRobot := struct {
		name string
		grabbers int
	} {
		name: "Robo",
		grabbers: 5,
	}

	fmt.Printf("Robot grabbers: %v \n", myRobot.grabbers)

	userMK := credentialsToStore{}
	userMK.name = "MK"
	userMK.location = "Waithaka"
	userMK.number = 71
	userMK.password = "9080$"
	fmt.Printf("Name making use of embedded structs: %v \n", userMK.name)

	userJG := credentialsToStore {
		password: "3758",
		location: "Githunguri",
		// Sort of nested syntax used when instantiating embedded struct for the first time
		user: user{
			name: "JG",
			number: 72,
		},
	}
	fmt.Printf("Name making use of embedded structs: %v \n", userJG.name)

	yayaCarSales := sales_data{}
	yayaCarSales.initial_stock = 100
	yayaCarSales.added_stock = 654
	yayaCarSales.sold_stock = 357
	yayaCarSales.remaining_stock = 248

	fmt.Printf("Discrepancy in car sales %v \n", yayaCarSales.stockDiscrepancy())

	mkInfo := authenticationInfo{}

	mkInfo.username = "MK"
	mkInfo.password = "hello"

	fmt.Println(mkInfo.getBasicAuth())



}