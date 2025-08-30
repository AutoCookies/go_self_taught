package main

import (
	"fmt"
)

func main() {
	var (
		a int = 42
		c int = 42
	)

	var (
		b int = 42
	)

	fmt.Println(a == b)
	fmt.Println(a == c)
}
