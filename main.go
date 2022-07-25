package main

import (
	"Creating-Custom-Data-Types-with-Go/organization"
	"fmt"
)

func main() {
	p := organization.NewPerson(
		"James",
		"Wilson",
		organization.NewEuropeanUnionIdentifier("123-45-6789", "Germany"),
	)
	err := p.SetTwitterHandler("@jam_wills")
	fmt.Printf("%T\n", p.TwitterHandler())

	if err != nil {
		fmt.Printf("An error occured setting twitter handler: %s\n", err.Error())
	}

	//name1 := Name{First: "James", Last: "Wilson"}
	//name2 := Name{First: "James2", Last: "Wilson"}

	ssn := organization.NewSocialSecurityNumber("123-45-6789")
	eu := organization.NewEuropeanUnionIdentifier("12345", "France")

	if ssn == eu {
		println("We match")
	}

	fmt.Printf("%T\n", ssn)
	fmt.Printf("%T\n", eu)
}

type Name struct {
	First string
	Last  string
}

type OtherName struct {
	First string
	Last  string
}
