package main

import "fmt"

type contactInfo struct {
	email string
	zip int
}

type person struct {
	firstName string
	lastName string
	contact contactInfo
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
		contact: contactInfo{
			email: "jim@gmail.com", 
			zip: 51952,
		},
	}

	fmt.Printf("\n%+v", jim)
}