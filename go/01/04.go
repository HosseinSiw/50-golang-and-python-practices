package main

import (
	"unicode"
)

/*
Write a program that accepts a sentence and calculate the number of upper case letters and lower case letters.
Suppose the following input is supplied to the program:
Hello world!
Then, the output should be:
UPPER CASE 1
LOWER CASE 8
*/

func Ex4(s string) map[string]int {
	var result = map[string]int{
		"Lower":        0,
		"Upper":        0,
		"White spaces": 0,
	}
	for i := 0; i < len(s); i++ {
		if unicode.IsLower(rune(s[i])) {
			result["Lower"] += 1

		} else if unicode.IsUpper(rune(s[i])) {
			result["Upper"] += 1

		} else {
			result["White spaces"] += 1
		}
	}
	return result
}
