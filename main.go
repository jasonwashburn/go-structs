package main

import "fmt"

type contactInfo struct {
	emailAddress string
	zip          int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			emailAddress: "jim@gmail.com",
			zip:          91000,
		},
	}

	jim.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
