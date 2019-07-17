package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

// implicitly names field name contactInfo with type contactInfo
// type person struct {
// 	firstName string
// 	lastName  string
// 	contactInfo
// }

func main() {
	// 1 of 3 to Declaring Struct
	// by order of string value it implicitly assigns the first and last name
	// alex := person{"Alex", "Anderson"}

	// 2 of 3 to Declaring Struct
	// explicit way (recommended)
	// alex := person{firstName: "Alex", lastName: "Anderson"}
	// fmt.Println(alex)

	// 3 of 3 to Declaring Struct
	// var alex person (will assign a zero value for string its an empty "")
	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}
	// &jim (&variableName) will give the memory address of the value this variable is pointing at
	// explicitly define the pointer before updating the value
	// jimPointer := &jim
	// jimPointer.updateName("jimmy")

	// implicitly uses pointer due to function signature
	jim.updateName("jimmy")
	jim.print()
}

// *person (type description) means we are working with a type pointer that points to a person.
func (pointerToPerson *person) updateName(newFirstName string) {
	// *pointer (gives the value this memory address is pointing at)
	// *pointerToPerosn (operator) means we want to manipulate the value the pointer is referencing
	(*pointerToPerson).firstName = newFirstName
}

// receiver function for person
// If you arent using the reciever value in the function in this case "p" then you can omit the value
// so it would be func (person) print() {}
func (p person) print() {
	fmt.Printf("%+v", p)
}
