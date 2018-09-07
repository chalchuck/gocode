package main

import (
	"fmt"
)

type contactInfo struct {
	emailAddress string
	mobileNumber string
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	charles := person{
		firstName: "Charles",
		lastName:  "Walker",
		contact: contactInfo{
			emailAddress: "charles@walker.app",
			mobileNumber: "0730333332",
		},
	}
	charles.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
