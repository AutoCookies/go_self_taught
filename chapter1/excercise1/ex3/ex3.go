package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	// Version 1: Using +=
	start := time.Now()
	s, sep := "", ""
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	fmt.Println("+= elapsed:", time.Since(start))

	// Version 2: Using strings.Join
	start = time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println("Join elapsed:", time.Since(start))
}
