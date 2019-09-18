package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
	contact   contactInfo
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
		contact: contactInfo{
			email:   "dand55@msn.com",
			zipCode: 55434,
		},
	}

	richie := person{"Richie", "Donovan", 25, contactInfo{"richie.donovan62@gmail.com", 55447}}

	var karen person
	karen.firstName = "Karen"
	karen.lastName = "Donovan"
	karen.age = 60
	karen.contact.email = "kdonovan3@msn.com"

	fmt.Println(dan)
	fmt.Println(richie)
	fmt.Printf("%+v", karen) // Prints out fields and values
}
