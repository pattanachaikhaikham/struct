package main

import "fmt"

type student struct {
	name  string
	age   int
	email string
}

func main() {
	var one student
	one.name = "Goku"
	one.age = 18
	one.email = "Goku@super.saiya"

	two := student{"Gohan", 3, "Gohan@super.saiya"}

	three := student{name: "Videl", email: "Videl@daughter.satan"}

	four := student{age: 20}

	fmt.Println(one)
	fmt.Println(two)
	fmt.Println(three)
	fmt.Println(four)
}
