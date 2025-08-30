package main

import (
	"fmt"
)

func main() {
	var student1 string = "who"
	var student2 string
	student1 += " am I"
	student1 = "Hello"
	student2 = student1

	fmt.Println(student1)
	fmt.Println(student2)
}
