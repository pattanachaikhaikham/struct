package main

type student struct {
	name  string
	age   int
	email string
}

func (std student) growUp(i int) {
	std.age = std.age + i
}

func main() {
	goku := student{name: "Goku"}
	pup := pupil{std: goku}
	pup.std.introduce()
}
