package main

import (
	"fmt"
	"os"
)

func main() {
	i := len(os.Args)

	for j := 0; j < i; j++ {
		fmt.Println(j, os.Args[j])
	}
}
