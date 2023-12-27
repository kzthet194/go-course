package main

import "log"

func main() {
	var myString string
	myString = "Green"

	log.Println("My String is set to", myString)

	changUsingPointer((&myString))
	log.Println("after function call myString is set to", myString)
}

func changUsingPointer(s *string) {
	log.Println("s is set to", s)
	newValue := "Red"
	*s = newValue
}
