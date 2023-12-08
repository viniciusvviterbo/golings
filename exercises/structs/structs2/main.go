// structs2
// Make me compile!
//
package main

import "fmt"

type Phone struct {
	countryCode string
	stateCode   string
	phoneNumber string
}

func (p Phone) String() string {
	return fmt.Sprintf("+%s (%s) %s", p.countryCode, p.stateCode, p.phoneNumber)
}

type Person struct {
	// don't just create the phone field here. embed a new struct
	phone Phone
	name  string
	age   int
}

func main() {
	// contactDetails := ContactDetails{}
	person := Person{name: "John", age: 32, phone: Phone{countryCode: "55", stateCode: "21", phoneNumber: "9999-9999"}}
	fmt.Printf("%s is %d years old and his phone is %s\n", person.name, person.age, person.phone)
}
