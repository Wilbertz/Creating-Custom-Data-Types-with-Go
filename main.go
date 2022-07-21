package main

import (
	"Creating-Custom-Data-Types-with-Go/organization"
	"fmt"
)

func main() {
	p := organization.NewPerson(
		"James",
		"Wilson",
	)
	err := p.SetTwitterHandler("@jam wills")
	if err != nil {
		fmt.Printf("An error occured setting twitter handler: %s\n", err.Error())
	}
	println(p.FullName())
}
