package main

import (
	"fmt"
)

func main() {
	var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
	b := map[string]int{"Oslo": 1, "Bergen": 2, "Trondheim": 3, "Stavanger": 4}

	fmt.Printf("a\t%v\n", a)
	fmt.Printf("b\t%v\n", b)

	c := make(map[string]int)
	c["one"] = 1
	c["two"] = 2
	c["three"] = 3
	fmt.Printf("c\t%v\n", c)

	c["one"] = 11
	fmt.Printf("c\t%v\n", c)
	delete(c, "two")
	fmt.Printf("c\t%v\n", c)

	value, ok := c["three"]
	fmt.Printf("value: %v, ok: %v\n", value, ok)
	value, ok = c["two"]
	fmt.Printf("value: %v, ok: %v\n", value, ok)
}
