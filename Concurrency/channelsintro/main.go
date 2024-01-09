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

	waitGroupLoopTest()

}

func waitGroupLoopTest() {
    done := make(chan bool)
	var test sync.WaitGroup
    values := []string{"a", "b", "c"}
    for _, v := range values {
        test.Add(1)
		/* important otherwise each iteration of the loop uses the same instance of the variable v so each closure shares that single variable */
		v := v 
		go func() {
			defer test.Done()
            fmt.Println(v)
            done <- true
        }()
    }

    // wait for all goroutines to complete before exiting
    for _ = range values {
        <-done
    }
	test.Wait()
}

