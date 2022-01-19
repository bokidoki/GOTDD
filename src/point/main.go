package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
}

type Group struct {
	person Person
}

func main() {
	group := Group{person: Person{FirstName: "Na", LastName: "Li"}}
	changeNameDelegate(group.person, "Na")
	fmt.Println(group.person)
}

func changeNameDelegate(person Person, newLastName string) {
	changeName(&person, newLastName)
	fmt.Println(person)
}

func changeName(person *Person, newLastName string) {
	person.LastName = newLastName
}
