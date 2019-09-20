package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	dan := person{
		firstName: "Dan",
		lastName:  "Donovan",
		age:       66,
		contactInfo: contactInfo{
			email:   "dand55@msn.com",
			zipCode: 55434,
		},
	}

	richie := person{"Richie", "Donovan", 25, contactInfo{"richie.donovan62@gmail.com", 55447}}

	var karen person
	karen.firstName = "Karen"
	karen.lastName = "Donovan"
	karen.age = 60
	karen.contactInfo.email = "kdonovan3@msn.com"
	karen.contactInfo.zipCode = 55434

	// richiePointer := &richie
	// richiePointer.updateName("Richie")

	// Lines 37/38 does the same exact thing as line 40
	richie.updateName("Richie")

	dan.print()
	richie.print()
	fmt.Printf("%+v\n", karen) // Prints out fields and values
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}
