package main

import "fmt"

type contactInfo struct {
	email string
	zip   int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	// alex := person{"Alex", "Anderson"}

	// alex := person{
	// 	firstName: "Alex",
	// 	lastName:  "Anderson",
	// }

	var alex person
	alex.firstName = "Alex"
	alex.lastName = "Anderson"

	jim := person{
		firstName: "Jim",
		lastName:  "Gold",
		contact: contactInfo{
			email: "sth@sth.com",
			zip:   432,
		},
	}

	jimPointer := &jim
	jimPointer.updateName("jimmy")
	jim.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
