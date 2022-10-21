package main

import "fmt"

type contactInfo struct {
	email  string
	mobile string
}

type person struct {
	firstName string
	lastName  string
	age       int
	contact   contactInfo
}

func main() {
	gino := person{
		firstName: "Gino",
		lastName:  "Cubeddu",
		age:       26,
		contact: contactInfo{
			mobile: "0797600000",
			email:  "gino@gmail.com",
		},
	}
	gino.print()
	gino.updateName("Jimmy")
	gino.print()
}

// Using * for person to state that we want the caller to pass itself as a
// pointer. Since a struct is a value argument it won't be updated unless
// we make it have a pointer.
func (p *person) updateName(name string) {
	p.firstName = name
}

func (p person) print() {
	// %+v will print out the the struct in {key: value} syntax
	fmt.Printf("%+v\n", p)
}
