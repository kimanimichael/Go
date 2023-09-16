package main

import "fmt"

//Interface with getMessage() function
type message interface {
	getMessage() string
}

type empty interface {

}

type birthdayMessage struct {
	bmessage string
	recipient string
}

//getMessage method with birthdayMessage as a receiver hence the type implements the interface
func (bmess birthdayMessage) getMessage() string {
	bmess.bmessage = fmt.Sprintf("Happy Birthday %s\n", bmess.recipient)

	return bmess.bmessage
}

type dateMessage struct {
	dmessage string
	recepient string
	date string
}

func (dmess dateMessage) getMessage() string {
	dmess.dmessage = fmt.Sprintf("Hello %s. Can't do Sunday. Let's do the %sth\n", dmess.recepient, dmess.date)
	return dmess.dmessage
}

type worker interface {
	getName() string
	getSalary() int
}

type contractor struct {
	name string
	hourlyPay int
	hoursPerYear int
}

type employee struct {
	name string
	monthlyPay int
}

func (em employee) getName() string {
	return em.name
}

func (em employee) getSalary() int {
	pay := em.monthlyPay * 12
	return pay
}

func (c contractor) getName() string {
	return c.name
}

func (c contractor) getSalary() int {
	pay := c.hourlyPay * c.hoursPerYear
	return pay
}

func printMessage(m message) {
	fmt.Printf(m.getMessage())
}

func getworkerDetails(e worker) {
	fmt.Printf("Name: %s \n", e.getName())
	fmt.Printf("Pay: %d \n", e.getSalary())
}

func main() {
	mkMessage := birthdayMessage{}
	mkMessage.recipient = "MK"
	printMessage(mkMessage)
	mwMessage := dateMessage{}
	mwMessage.recepient = "Waringa"
	mwMessage.date = "4"
	printMessage(mwMessage)

	contractor01 := contractor{}
	contractor01.name = "Renny"
	contractor01.hourlyPay = 1000
	contractor01.hoursPerYear = 2000

	employee01 := employee{}
	employee01.name = "SM"
	employee01.monthlyPay = 90000

	getworkerDetails(contractor01)
	getworkerDetails(employee01)

}