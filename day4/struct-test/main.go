package main

import "fmt"

type Person struct {
	Name    string
	Age     int
	Phone   string
	Address string
	Email   string
}

func main() {
	var person1 Person
	person1.Name = "ALice"
	person1.Age = 20
	person1.Phone = "123456789"
	person1.Address = "123 Main St"
	person1.Email = "xgTQ0@example.com"
	fmt.Println(person1.Name, person1.Age, person1.Phone, person1.Address, person1.Email)
	fmt.Println(person1)
}
