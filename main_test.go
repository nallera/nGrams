package main

import "testing"

func TestPunctuationStrip(t *testing.T) {
	stringInput := "t,,his is j,;ust. .a t..est, cas/e.,"

	cleanString := stripRegex(stringInput)

	if cleanString != "this is just a test case" {
		t.Errorf("Wrong cleanString, should be \"this is just a test case\" but instead is %s", cleanString)
	}
}

func TestNgramForEmptyString(t *testing.T) {
	stringInput := ""

	nGrams := generateNGrams(stringInput)

	if len(nGrams) > 0 {
		t.Errorf("Empty string should generate no nGrams, but got %v", nGrams)
	}
}

func TestNgramForSingleWord(t *testing.T) {
	stringInput := "test"
	want := []string{"test"}

	nGrams := generateNGrams(stringInput)

	if len(nGrams) != 1 {
		t.Errorf("Want nGrams=%v, but got %v", want, nGrams)
	}
	if nGrams[0] != want[0] {
		t.Errorf("Difference in only nGram: \"%s\" but want \"%s\"", nGrams[0], want[0])
	}
}

func TestNgramForThreeWords(t *testing.T) {
	stringInput := "simple test case"
	want := []string{"simple", "simple test", "simple test case", "test", "test case", "case"}

	nGrams := generateNGrams(stringInput)

	for i, nGram := range nGrams {
		if nGram != want[i] {
			t.Errorf("Difference in nGram number %d: \"%s\" but want \"%s\"", i, nGram, want[i])
		}
	}
}

func TestExtraSpacesAreStripped(t *testing.T) {
	stringInput := "simple     test                   case"
	want := []string{"simple", "simple test", "simple test case", "test", "test case", "case"}

	nGrams := generateNGrams(stringInput)

	for i, nGram := range nGrams {
		if nGram != want[i] {
			t.Errorf("Difference in nGram number %d: \"%s\" but want \"%s\"", i, nGram, want[i])
		}
	}
}
