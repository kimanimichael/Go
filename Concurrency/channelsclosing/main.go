package main

import (
	"fmt"
	"time"
)

func main() {
	test(3)
}

func countReports(numSentCh chan int) int {
	total := 0

	for {
		numSent, ok := <-numSentCh

		if !ok {
			break
		}
		total += numSent
	}

	return total
}

func test(numBatches int) {
	numSentCh := make(chan int)
	go sendReports(numBatches, numSentCh)

	fmt.Println("Start counting")
	numReports := countReports(numSentCh)

	fmt.Printf("%d reports sent \n",numReports)
	fmt.Println("=========================================")
	
}

func sendReports(numBatches int, ch chan int) {
	for i := 0; i < numBatches; i++ {
		numReports := i*32 + 32%17
		ch <- numReports
		fmt.Printf("Sent a bunch of %d reports\n", numReports)
		time.Sleep(time.Millisecond * 100)
	}
	close(ch)
}