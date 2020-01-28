package main

type I interface{}

func main() {
	var i I
	i = 3.14159265359
	desc(i)
	i = "Hello World"
	desc(i)
}
