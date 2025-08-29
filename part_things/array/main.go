package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Scanf("%d", &a)

	var arr1 = [...]int{1, 2, 3, 4, 5, 6}
	arr2 := [5]int{4, 5, 6, 7, 8}

	fmt.Println(a)
	fmt.Println(arr1)
	fmt.Println(arr2)
}
