package main

import (
	"Creating-Custom-Data-Types-with-Go/organization"
	"fmt"
)

func main() {
	p := organization.NewPerson(
		"James",
		"Wilson",
		organization.NewSocialSecurityNumber("123-45-6789"),
	)
	err := p.SetTwitterHandler("@jam_wills")
	fmt.Printf("%T\n", p.TwitterHandler())

	if err != nil {
		fmt.Printf("An error occured setting twitter handler: %s\n", err.Error())
	}
	println(p.TwitterHandler())
	println(p.TwitterHandler().RedirectUrl())
	println(p.ID())
	println(p.Name.FullName())
}
