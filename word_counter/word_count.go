package main

import (
	"flag"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	file := flag.String("file", "", "a text file")
	flag.Parse()

	text, err := os.ReadFile(*file)
	if err != nil {
		fmt.Println(err)
	}

	words := strings.Fields(string(text))

	for k, v := range words {
		words[k] = strings.ToLower(strings.Trim(v, "~!@#$%^&*()-_=+[{]}\\|;:,<.>/?"))
	}

	slices.Sort(words)

	wordCount := make(map[string]int)

	for _, v := range words {
		wordCount[v]++
	}

	uniqueWords := make([]string, 0, len(words))
	for i, w := range words {
		if i == 0 || w != words[i-1] {
			uniqueWords = append(uniqueWords, w)
		}
	}

	for _, v := range uniqueWords {
		fmt.Printf("%s: %d\n", v, wordCount[v])
	}

}
