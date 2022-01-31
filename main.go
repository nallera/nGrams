package main

import (
	"flag"
	"fmt"
	"regexp"
	"strings"
)

func main() {
	stringInput := flag.String("s", "this is an example", "the string from which to generate the n-grams")

	flag.Parse()

	cleanString := stripRegex(*stringInput)
	print(fmt.Sprintf("cleanString: %v\n", cleanString))

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