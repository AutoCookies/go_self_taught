package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	imput := bufio.NewScanner(os.Stdin)
	for imput.Scan() {
		counts[imput.Text()]++
	}
	// Note: ignoring potential errors from imput.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
