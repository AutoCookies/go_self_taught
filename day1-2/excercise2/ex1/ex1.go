package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	files := os.Args[1:] // skip program name
	for _, filename := range files {
		f, err := os.Open(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup: %v\n", err)
			continue
		}

		scanner := bufio.NewScanner(f)
		lineCounts := make(map[string]int)
		hasDuplicate := false

		for scanner.Scan() {
			line := scanner.Text()
			lineCounts[line]++
			if lineCounts[line] > 1 {
				hasDuplicate = true
				break
			}
		}

		if err := scanner.Err(); err != nil {
			fmt.Fprintf(os.Stderr, "reading %s: %v\n", filename, err)
		}

		if hasDuplicate {
			fmt.Println(filename)
		}

		f.Close()
	}
}
