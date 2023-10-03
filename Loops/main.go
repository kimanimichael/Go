package main

import (
	"fmt"
)

func bulkSend(numMessages  int) float64 {
	baseCost := 1
	additionalCost := 0.01
	var totalCost float64
	for i := 0; i < numMessages; i++ {
		// cost := float64(baseCost * (i+1)) + (0.01* float64(i))
		totalCost += float64(baseCost) + (additionalCost* float64(i))
	}
	return totalCost
}

func main () {
	fmt.Printf("Total cost %.2f \n", bulkSend(10))

}