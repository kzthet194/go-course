package main

import "fmt"

func main() {
	fmt.Println("Hello, World.")

	var whatToSay string
	var i int

	whatToSay = "Hello, cruel world"
	fmt.Println(whatToSay)

	i = 7
	fmt.Println("i is set to", i)

	whatWasSaid, theOtherThingThatWasSaid := saySomething()
	fmt.Println("the function returned", whatWasSaid, theOtherThingThatWasSaid)
}

func saySomething() (string, string) {
	return "something", "else"
}
