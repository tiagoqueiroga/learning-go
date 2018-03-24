package main

import "fmt"

type contactInfo struct {
	email   string
	zipcode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	alex := person{
		firstName: "Alex",
		lastName:  "Jones",
		contactInfo: contactInfo{
			email:   "alex.jones@gmail.com",
			zipcode: 99999,
		},
	}

	alex.updateName("John")
	fmt.Println(alex)
}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

func (p person) print() {
	fmt.Println(p)
}
