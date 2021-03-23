package main

import (
	"fmt"
	"strings"
)

func main() {

	pattern := "abba"
	str := "dog cat cat dog"
	fmt.Println(wordPattern(pattern, str))
}

func wordPattern(pattern string, s string) bool {

	word2ch := map[string]byte{}
	ch2word := map[byte]string{}
	wordList := strings.Split(s, " ")
	if len(pattern) != len(wordList) {
		return false
	}
	for i, word := range wordList {
		ch := pattern[i]
		if word2ch[word] > 0 && word2ch[word] != ch || ch2word[ch] != "" && ch2word[ch] != word {
			return false
		}
		word2ch[word] = ch
		ch2word[ch] = word
	}
	return true
}
