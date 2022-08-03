package highestscoreword

import "strings"

/*
Given a string of words, you need to find the highest scoring word.
Each letter of a word scores points according to its position in the alphabet: a = 1, b = 2, c = 3 etc.
You need to return the highest scoring word as a string.
If two words score the same, return the word that appears earliest in the original string.
All letters will be lowercase and all inputs will be valid.
*/

func High(s string) string {
	var highest string
	var maxSum int
	for _, word := range strings.Split(s, " ") {
		var sum int
		for _, char := range word {
			sum += int(char) - ('a' - 1)
		}
		if sum > maxSum {
			maxSum = sum
			highest = word
		}
	}
	return highest
}
