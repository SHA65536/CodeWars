package wheremyanagrams

/*
What is an anagram? Well, two words are anagrams of each other if they both contain the same letters.
Write a function that will find all the anagrams of a word from a list.
You will be given two inputs a word and an array with words.
You should return an array of all the anagrams or an empty array if there are none.
*/

func Anagrams(word string, words []string) []string {
	var res = []string{}
	var freqs = map[byte]int{}
	for idx := range word {
		freqs[word[idx]]++
	}
EachWord:
	for _, curword := range words {
		if len(word) != len(curword) {
			continue
		}
		var cur = map[byte]int{}
		for idx := range curword {
			cur[curword[idx]]++
		}
		for key, val := range freqs {
			if cur[key] != val {
				continue EachWord
			}
		}
		res = append(res, curword)
	}
	if len(res) == 0 {
		return nil
	}
	return res
}
