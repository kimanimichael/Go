package main

import "fmt"

type expense interface {
	cost() float64
}

type printer interface {
	print()
}

type email struct {
	isSubscribed bool
	body string
	toAddress string
}

type sms struct {
	body string
	toPhoneNUmber string
	isSubscribed bool
}

func (em email) cost() float64{
	if !em.isSubscribed {
		cost := 0.05 * float64(len(em.body))
		return cost
	}
	return 0.01 * float64(len(em.body))

}

func (s sms) cost() float64 {
	if !s.isSubscribed {
		cost := 0.1 * float64(len(s.body))
		return cost
	}
	cost := 0.03 * float64(len(s.body))
	return cost 
}

func (em email) print() {
	fmt.Printf("%s \n",em.body)
}

func getEmailDetails(e expense, p printer) {
	p.print();
	fmt.Printf("Cost of the email is %f \n", e.cost())
}

//Type assertions
func getExpenseReport (e expense) (string, float64) {
	ex, ok := e.(email)
	if ok {
		return ex.toAddress, ex.cost()
	}
	sm, ok := e.(sms)
	if ok {
		return sm.toPhoneNUmber, sm.cost()
	}
	return "", 0.0
}

func test (e expense) {
	recepient, cost := getExpenseReport(e)
	
	switch e.(type) {
	case email:
		fmt.Printf("Report: The email will be sent to %s and will cost %f \n", recepient, cost)
	case sms:
		fmt.Printf("Report: The sms will be sent to %s and will cost %f \n", recepient, cost)
	default:
		fmt.Printf("Report: Invalid expense \n")
	}
}

func main() {
	leaveEmail := email{}
	leaveEmail.isSubscribed = true
	leaveEmail.body = "Greetings. I wish to apply for leave on the 22nd."
	leaveEmail.toAddress = "mk@novek.io"
	getEmailDetails(leaveEmail, leaveEmail)

	leaveSms := sms{}
	leaveSms.body = "Hello man. I'll be on leave on the 22nd"
	leaveSms.isSubscribed = true
	leaveSms.toPhoneNUmber = "0724612024"
	test(leaveEmail)
	test(leaveSms)
}