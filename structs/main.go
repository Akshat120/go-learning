package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	// contact contactInfo
	contactInfo
}

func main() {
	// alex := person{
	// 	firstName: "Alex",
	// 	lastName:  "Anderson",
	// }

	// var alex person
	// fmt.Printf("%+v\n", alex)
	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"
	// alex.contact = contactInfo{"alex@gmail.com", 123456}
	// fmt.Printf("%+v\n", alex)

	jim := person{
		firstName: "Jim",
		lastName:  "Pearson",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 123456,
		},
	}

	// jimPointer := &jim
	// jimPointer.updateFirstName("jimmy")
	jim.updateFirstName("jimmy")
	jim.print()
}

func (pointerToPerson *person) updateFirstName(firstName string) {
	(*pointerToPerson).firstName = firstName
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}
