package main

import (
	"fmt"
	"sync"
)

func main() {
	checkEven(10)
}


func checkEven(num int) {
	/* create channel */
	isEvenChan := make(chan bool)

	/* create wait group - allows waiting so all go routines can finish before exiting the program */
	var even sync.WaitGroup

	go func ()  {
		/* adds 1 to the even wait group */
		even.Add(1)
		for i:=0; i < num; i++ {
			if i % 2 == 0 {
				isEvenChan <- true
				continue
			}
			isEvenChan <- false
	}
	/* subtracts one from the even wait group*/
	even.Done()
	}()

	for i:=0; i < num; i++ {
		isEven := <- isEvenChan
		fmt.Println("Num is even:", isEven)
	}
	/* waits until the even wait group is zero before exiting the function*/
	even.Wait()
	
}

