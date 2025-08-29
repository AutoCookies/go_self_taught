package main

import (
	"fmt"
)

func main() {
	c1 := complex(1, 2)
	c2 := 2i + 4

	fmt.Println(c1)
	fmt.Println(c2)
	fmt.Println(c1 + c2)
}
