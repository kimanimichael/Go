package main

import (
	"fmt"
	"sync"
)

func main() {
	checkEven(10)
}

var even sync.WaitGroup

func checkEven(num int) {
	isEvenChan := make(chan bool)
	
	
	go func ()  {
		even.Add(1)
		for i:=0; i < num; i++ {
			if i % 2 == 0 {
				isEvenChan <- true
				continue
			}
			isEvenChan <- false
	}	
	even.Done()
	}()
	isEven := <- isEvenChan
	fmt.Println("Num is even:", isEven)
	even.Wait()
	
}

