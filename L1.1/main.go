package main

import "fmt"

type Human struct {
	age  int
	name string
}

func (h *Human) information() {
	fmt.Printf("Name: %s\nage: %d\n", h.name, h.age)
}

type Action struct {
	Human
}

func (a Action) information() {
	a.Human.information()
}

func main() {
	person := Human{
		age:  30,
		name: "Стив",
	}
	person.information()

	actionPerson := Action{
		Human: Human{
			age:  20,
			name: "name",
		},
	}
	actionPerson.information()
}
