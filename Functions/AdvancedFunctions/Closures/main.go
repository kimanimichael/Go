package main

import "fmt"

func main() {
	sentenceFormer := concatter()
	sentenceFormer("There")
	sentenceFormer("goes")
	sentenceFormer("Michael")
	sentenceFormer("Kimani.")
	sentenceFormer("The")
	sentenceFormer("best")
	sentenceFormer("to")
	sentenceFormer("ever")
	sentenceFormer("do")
	sentenceFormer("it.")

	fmt.Println(sentenceFormer(""))
}

func concatter()  func(string) string {
	doc := ""
	/* this function references and assigns to the doc variable which is outside its own function body */
	return func(word string) string {
		doc += word + " "
		return doc
	}
}