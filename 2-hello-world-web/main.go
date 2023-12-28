package main

import "log"

type User struct {
	FirstName string
	LastName  string
}

func main() {
	myMap := make(map[string]User)

	me := User{
		FirstName: "Ko Zin",
		LastName:  "Thet",
	}

	myMap["me"] = me

	log.Println(myMap["me"].FirstName, myMap["me"].LastName)

	name := []string{"one", "two", "three"}

	log.Println(name)
}
