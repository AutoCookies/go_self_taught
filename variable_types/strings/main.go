package main

import (
	"fmt"
)

type Fate int

func sum(a int, b int) int {
	return a + b
}

func main() {
	name := "AutoCookies"
	fmt.Printf("My name is %s\n", name, "With type: ", fmt.Sprintf("%T", name))

	const (
		Decided Fate = iota
		Undecided
		Unknown
	)

	fmt.Println(Decided, Undecided, Unknown)
	fmt.Println("Sum is: ", sum(5, 10))
}
