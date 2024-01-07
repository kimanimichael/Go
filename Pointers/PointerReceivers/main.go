package main

import "fmt"

type car struct {
	colour string
}
/* function with a value receiver */
func (c car) setColour(colour string) {
	c.colour = colour
}
/* function with a pointer receiver */
func (c *car) setColourPointer(colour string) {
	c.colour = colour
}

func main() {
	probox := car{
		colour: "white",
	}
	fmt.Println(probox.colour)
	/* function with value reveiver does not mutate the original probox colour*/
	probox.setColour("Blue")
	fmt.Println(probox.colour)

	/* function with pointer reveiver does mutate the original probox colour*/
	/* this call is allowed as the value is cast under the hood to a pointer sort of*/
	probox.setColourPointer("Blue")
	fmt.Println(probox.colour)
}