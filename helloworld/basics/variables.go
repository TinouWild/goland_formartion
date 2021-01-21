package basics

import (
	"fmt"
)

var Name = "Etienne"
var password = "secretpassword"

func TypeOfVariables() {
	var age int = 20
	fmt.Println(age)
	fmt.Println(password)

	var weight, height int = 80, 190
	fmt.Println(weight, height)

	var lastname = "Bob"
	fmt.Println("inférence de type string avec", lastname)

	email := "bob@go.org"
	fmt.Println(":=, le : remplace le var, indique une déclaration", email)
	email = "alice@go.org"
	fmt.Println("réaffectation possible")
}