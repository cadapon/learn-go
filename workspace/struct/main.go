package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email string
	zip   int
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Stacks",
		contactInfo: contactInfo{
			email: "jimmystacks@gmail.com",
			zip:   91910,
		},
	}
	fmt.Printf("%+v", jim)
}
