package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 4, 5}
	arr1 := [6]int{1, 2, 3, 4, 5, 6}
	slice := arr1[1:4]

	myslice := make([]int, 5, 10)

	fmt.Println(nums)
	fmt.Println(len(nums))
	fmt.Println(cap(nums))
	fmt.Printf("Slice = %v\n", slice)
	fmt.Println(len(myslice))
	fmt.Println(cap(myslice))
	fmt.Println(myslice)

	myslice = append(myslice, 1, 2, 3)
	fmt.Println(myslice)
	myslice[0] = 100
	fmt.Println(myslice)

	slice3 := append(myslice, slice...)
	fmt.Println(slice3)
}
