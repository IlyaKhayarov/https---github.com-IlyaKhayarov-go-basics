package main

import (
	"fmt"
)

type Person struct {
	Name    string
	Surname string
	Age     int
}
type pers func(*Person)

func main() {
	person1 := NewPerson()
	person2 := NewPerson(Name("Name2"))
	person3 := NewPerson(Name("Name3"), Surname("Surname3"))
	fmt.Printf("Person1 = %v\nPerson2 = %v\nPerson3 = %v", person1, person2, person3)
}

func NewPerson(p ...pers) *Person {
	i := &Person{
		Name:    "Name1",
		Surname: "Surname1",
		Age:     30,
	}
	for _, v := range p {
		v(i)
	}
	return i
}
func Name(p1 string) pers {
	return func(i *Person) {
		i.Name = p1
	}
}
func Surname(p2 string) pers {
	return func(i *Person) {
		i.Surname = p2
	}
}
