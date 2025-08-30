package main

import (
	"fmt"
)

func main() {
	// Example 1: Simple if-else statement, two slice can not be compared directly
	/*
		nums1 := []int{1, 2, 3, 4, 5}
		nums2 := []int{1, 2, 3, 4, 5}

		if nums1 == nums2 {
			fmt.Println("The two slices are equal.")
		} else {
			fmt.Println("The two slices are not equal.")
		}
	*/
	// Example 2: Using reflect.DeepEqual to compare two slices
	/*
		nums1 := []int{1, 2, 3, 4, 5}
		nums2 := []int{1, 2, 3, 4, 5}
		if equal := reflect.DeepEqual(nums1, nums2); equal {
			fmt.Println("The two slices are equal.")
		} else {
			fmt.Println("The two slices are not equal.")
		}
	*/
	// Example 3: Compare two array
	nums1 := [5]int{1, 2, 3, 4, 5}
	nums2 := [5]int{1, 2, 3, 4, 5}
	if nums1 == nums2 {
		fmt.Println("The two arrays are equal.")
	} else {
		fmt.Println("The two arrays are not equal.")
	}
}
