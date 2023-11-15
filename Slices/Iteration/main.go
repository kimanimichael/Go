package main

import "fmt"

func indexOfFirstBadWord(msg []string, badWords []string) int{
	for i, word := range msg {
		/* ignore index in this syntactic sugar iteration */
		for _, badWord := range badWords {
			if word == badWord {
				return i
			}
		}
	}
	return -1
}

func main() {
	stationery := []string{"pencil", "bluePen", "redPen", "blackPen"}

	/* syntatctic sugar for iteration */
	for i, item := range stationery {
		fmt.Println(i, item)
	}

	badWords := []string{"fuckup", "shit", "prick", "dude", "fuckery"}

	message := []string {"You", "absolute", "piece", "of", "shit", "fuckup"}
	message2 := []string {"You", "absolute", "beauty"}
	
	fmt.Println(indexOfFirstBadWord(message, badWords))
	fmt.Println(indexOfFirstBadWord(message2, badWords))
}