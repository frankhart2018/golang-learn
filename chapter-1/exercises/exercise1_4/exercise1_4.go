// Modify dup2 to print the names of all files in which each duplicated line occurs
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	filesMap := make(map[string][]string)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts, filesMap)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, filesMap)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%v\n", n, line, filesMap[line])
		}
	}
}

// Changes to counts map will be reflected in the original map (call by reference)
func countLines(f *os.File, counts map[string]int, filesMap map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		counts[line]++
		filesMap[line] = append(filesMap[line], f.Name())
	}
}
