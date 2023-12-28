package main

import "log"

func main() {
	for i := 0; i <= 10; i++ {
		log.Println(i)
	}

	animals := []string{"dog", "cat", "fist", "snake", "cow"}

	for i, animals := range animals {
		log.Println(i, animals)
	}

	for _, animals := range animals {
		log.Println(animals)
	}

	animals1 := make(map[string]string)

	animals1["dog"] = "Fido"
	animals1["cat"] = "Fluffy"

	for animalType, animalName := range animals1 {
		log.Println(animalType, animalName)
	}

	type User struct {
		FirstName string
		LastName  string
		Email     string
		Age       int
	}

	var users []User
	users = append(users, User{"Ko Zin", "Thet", "zin@gmail.com", 20})
	users = append(users, User{"Mg", "Mg", "mg@gmail.com", 23})
	users = append(users, User{"Hla", "Hla", "hla@gmail.com", 24})
	users = append(users, User{"Mya", "Mya", "mya@gmail.com", 21})

	for _, u := range users {
		log.Println(u.FirstName, u.LastName, u.Email, u.Age)
	}
}
