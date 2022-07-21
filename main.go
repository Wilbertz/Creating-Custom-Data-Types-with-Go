package main

import "Creating-Custom-Data-Types-with-Go/organization"

func main() {
	p := organization.NewPerson(
		"James",
		"Wilson",
	)
	println(p.ID())
	println(p.FullName())
}
