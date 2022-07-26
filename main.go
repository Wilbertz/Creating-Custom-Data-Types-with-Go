package main

import (
	"Creating-Custom-Data-Types-with-Go/organization"
	"fmt"
)

func main() {
	p := organization.NewPerson(
		"James",
		"Wilson",
		organization.NewEuropeanUnionIdentifier(123, "Germany"),
	)
	err := p.SetTwitterHandler("@jam_wills")
	fmt.Printf("%T\n", p.TwitterHandler())

	if err != nil {
		fmt.Printf("An error occured setting twitter handler: %s\n", err.Error())
	}
	println(p.ID())
	println(p.Country())
}
