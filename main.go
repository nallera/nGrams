package main

import (
	"flag"
	"regexp"
	"strings"
)

// To install and run the project simply save both files to the same directory, and run 'go build main.go' in the console.
// Then run the main executable that should've been created: 'main -s <string>'. <string> is the text from which to create the n-grams.
// To run the tests run 'go mod init', 'go mod tidy' and 'go test -v' in the console.

func main() {
	stringInput := flag.String("s", "this is an example", "the string from which to generate the n-grams")
	flag.Parse()

	cleanString := stripRegex(*stringInput)

	nGrams := generateNGrams(cleanString)

	for _,nGram := range nGrams {
		println(nGram)
	}
}

func generateNGrams(s string) []string {
	var initial [][]string
	var nGrams []string

	words := strings.Fields(s)

	for end := 0; end < len(words); end++ {
		initial = append(initial, words[0:end+1])
	}

	for start := 0; start < len(words); start++ {
		for end := start; end < len(words); end++ {
			nGrams = append(nGrams, strings.Join(initial[start][start:end+1], " "))
		}
	}

	return nGrams
}

func stripRegex(in string) string {
	reg, _ := regexp.Compile("[^a-zA-Z0-9 ]+")
	return reg.ReplaceAllString(in, "")
}