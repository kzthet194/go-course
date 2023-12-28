package main

import "log"

type myString struct {
	FirstName string
}

func (s myString) printFirstName() string {
	return s.FirstName
}

func main() {
	var myVar myString
	myVar.FirstName = "First"

	myVar2 := myString{
		FirstName: "Second",
	}

	log.Println("myVar is set to", myVar.FirstName)
	log.Println("myVar2 is set to", myVar2.FirstName)

	log.Println("myVar is set to", myVar.printFirstName())
	log.Println("myVar2 is set to", myVar2.printFirstName())
}
