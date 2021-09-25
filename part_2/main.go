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
		lastName:  "Anderson",
		contactInfo: contactInfo{
			email:   "will.a@cmcst.com",
			zipcode: 45235,
		},
	}
	alex.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
