package main

import (
	"fmt"
	// "time"
)

func main() {
	concurrentFib(10)
}

func concurrentFib(n int) {
	numCh := make(chan int)
	go func ()  {
		fibonacci2(numCh,n)
	}()
	for channelNumber := range numCh {
		fmt.Println(channelNumber)
	}
	
}

func fibonacci(fibCh chan int, num int) {
	var prevNum int
	var currentNum int
	var newNum int
	prevNum = 1
	fibCh <- prevNum
	currentNum = 2
	fibCh <- currentNum
	for i := 2; i < num; i++ {
		newNum = currentNum + prevNum
		prevNum = currentNum
		currentNum = newNum
		fibCh <- newNum
	}
	close(fibCh)
}

func fibonacci2(fibCh chan int, num int) {
	x, y := 0, 1
	for i := 0; i < num; i++ {
		fibCh <- x
		x, y = y, x + y
		// time.Sleep(time.Millisecond * 10)
	}
}
/*
1
2
3
5
8
13
21
34
55
89
*/