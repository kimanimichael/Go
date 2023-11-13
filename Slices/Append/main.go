package main

import (
	"fmt"
)

func createMessageSlice () []string {
	var messageStrings [3]string
	messageStrings[0] = "Click here to sign up"
	messageStrings[1] = "Pretty please click here"
	messageStrings[2] = "We demand that you sign up"

	messageSlice := messageStrings[:]
	return messageSlice
}

type cost struct {
	day int
	value float64
}



func getCostsByDay(costs []cost) []float64 {
	costByDay := []float64{}

	for i:=0; i<len(costs); i++ {
		cost := costs[i]
		for cost.day >= len(costByDay) {
			costByDay = append(costByDay, 0.0)
			
		}
		
		costByDay[cost.day] += cost.value

	}
	return costByDay
}

func main() {
	messageSlice := createMessageSlice()
	fmt.Println(messageSlice)
	/* test append function */
	messageSlice = append(messageSlice, "MK")
	fmt.Println(messageSlice)
	
	/* one variable of type cost*/
	holidayCost := cost{
		8, 7.8}

	fmt.Println(holidayCost.value)

	/* slices of type cost */
	weekCost:= []cost{
		{0, 4.0},
		{1, 2.1},
		{1, 3.1},
		{5, 2.5},
	}

	fmt.Println(len(weekCost))

	fmt.Println(getCostsByDay(weekCost))
}