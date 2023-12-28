package main

import "log"

func main() {
	cat := "cat"

	if cat == "cat" {
		log.Println("cat is cat")
	} else {
		log.Println("cat is not cat")
	}

	myNum := 100
	isTue := false

	if myNum > 99 && isTue {
		log.Println("myNum is greater than 99 and isTrue is True")
	} else if myNum > 99 && !isTue {
		log.Println("myNum is greater than 99 and isTrue is False")
	}

	myCat := "cat"

	switch myCat {
	case "cat":
		log.Println("myCat is cat")
	case "dog":
		log.Println("myCat is dog")
	case "fish":
		log.Println("myCat is fish")
	default:
		log.Println("myCat is something else")
	}
}
