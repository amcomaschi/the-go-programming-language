package main

import (
	"os"
	"fmt"
	"bufio"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]

	if len(files) != 0 {
		for _, arg := range files {

			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}

			input := bufio.NewScanner(f)

			for input.Scan() {
				counts[input.Text()]++
			}
			f.Close()

			for line, n := range counts {
				if n > 1 {
					fmt.Printf("%s %d\t%s\n", f.Name(), n, line)
				}
			}
		}
	}
}