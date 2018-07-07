package main

import (
	"fmt"
)

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	andrew := person{
		firstName: "Andrew",
		lastName:  "Smith",
		contact: contactInfo{
			email:   "andrew@smith.com",
			zipCode: 56000,
		},
	}

	andrew.updateName("andreas")
	andrew.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
