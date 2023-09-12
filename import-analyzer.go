package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func getImportPath(stm string) string {
	var path string

	fields := strings.Fields(stm)
	for i, f := range fields {
		if f == "from" {
			path = fields[i+1]
		}
	}

	if path == "" {
		return path
	}

	var found bool
	var r rune
	start := 0
	for !found && start < len(path) {
		r = rune(path[start])
		if !unicode.IsSpace(r) && r != '\'' && r != '"' {
			found = true
		} else {
			start++
		}
	}

	return path[start : strings.LastIndex(path, "/")+1]
}

func processStatement(paths *[]string, s string) {
	if strings.Contains(s, "import") {
		p := getImportPath(s)
		if p != "" {
			*paths = append(*paths, p)
		}
	}
}

func main() {
	inputFile := flag.String("f", "", "Name of the file to use for input")
	flag.Parse()

	var paths []string

	if *inputFile != "" {
		// TODO: read file for content
		fmt.Println("Provided file name: ", *inputFile)
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			s := scanner.Text()
			processStatement(&paths, s)
		}
	}

	fmt.Println("Paths:")

	occurrences := make(map[string]int)

	for _, p := range paths {
		_, ok := occurrences[p]
		if !ok {
			occurrences[p] = 0
		}
		occurrences[p] += 1
	}

	fmt.Println("Occurrences:")
	for k, v := range occurrences {
		fmt.Printf("[%s]: %d\n", k, v)
	}
}
