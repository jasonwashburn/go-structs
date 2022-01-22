package main

import "fmt"

type contactInfo struct {
	emailAddress string
	zip          int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			emailAddress: "jim@gmail.com",
			zip:          91000,
		},
	}

	fmt.Printf("%+v", jim)
}
