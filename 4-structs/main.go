package main

import "fmt"

type contactInfo struct {
	email string
	zip int
}

type person struct {
	firstName string
	lastName string
	contactInfo
}

func main() {
	alex := person{firstName: "alex", lastName: "anderson"} // Creating a new struct
	fmt.Println(alex)


	var alex2 person
	fmt.Printf("%#v", alex2)

	alex2.firstName = "John"
	alex2.lastName = "Hopper"
	fmt.Printf("\n%#v", alex2)

	jim := person {
		firstName: "Jim",
		lastName: "Hopper",
		contactInfo: contactInfo{
			email: "jim@gmail.com", 
			zip: 51952,
		},
	}

	jim.print()

	jimPtr := &jim
	jimPtr.updateName("Joseph", "Hopping")
	jim.print()
}

func (ptrToPerson *person) updateName(newFirst string, newLast string) {
	ptrToPerson.firstName = newFirst
	ptrToPerson.lastName = newLast
}

func (p person) print() {
	fmt.Printf("\n%+v", p)
}