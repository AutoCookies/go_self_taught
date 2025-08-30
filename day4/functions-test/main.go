package main

import (
	"fmt"
)

func printMySelf(name string, age int) {
	fmt.Printf("My name is %s and I am %d years old. \n", name, age)
}

func combineString(a string, b string) string {
	return a + b
}

func sums(a int) int {
	if a == 0 {
		return 0
	}
	return a + sums(a-1)
}

func main() {
	var name string = "AutoCookies"
	var age int = 25
	printMySelf(name, age)
	result := combineString("Hello, ", "World!")
	fmt.Println(result)
	fmt.Println(sums(5))
}
