package main

import (
	"fmt"
)

func main() {
	number := 3
	switch number {
	case 1, 3, 5, 7, 9:
		fmt.Println("Odd number")
	case 2, 4, 6, 8, 10:
		fmt.Println("Even number")
	default:
		fmt.Println("Other number")
	}
}
