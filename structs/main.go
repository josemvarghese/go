package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}
type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	// user := person{
	// 	firstName: "Alex",
	// 	lastName:  "Anderson",
	// }
	var user person
	user.firstName = "Alex"
	user.lastName = "Anderson"
	// fmt.Println(user)
	// fmt.Printf("%+v", user)

	jim := person{
		firstName: "jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jim@mail.com",
			zipCode: 12312,
		},
	}
	// jim.print()
	// jim.updateFirstName("Jose")
	// jim.print()
	/*
	   :::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
	   &jim  -  give the memory address of the value that variable(jim) pointing
	   :::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
	*/
	// jimPointer := &jim
	// fmt.Print(&jim)
	// jimPointer.updateFirstName("Jose")
	jim.updateFirstName("Jose")
	jim.print()
}
func (pointerToPerson *person) updateFirstName(newFirstName string) {
	/*
	   :::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
	   /*pointerToPerson  -  gives the value of that memory address is pointing
	   :::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
	*/
	(*pointerToPerson).firstName = newFirstName
	// p.print()
}
func (p person) print() {
	// p.firstName = "Jose"
	fmt.Printf("%+v", p)
}

/*
:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
Value types

int
float
string
bool
structs

Required pointers to change these things in functions
:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
*/

/*
:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
Reference  types

slices
maps
channels
pointers
functions

we can directily change the values
:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
*/
