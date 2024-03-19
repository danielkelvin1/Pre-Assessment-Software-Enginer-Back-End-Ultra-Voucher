package main

import (
	"fmt"
	"strings"
)

func main() {
	words := []string{"cook", "save", "taste", "aves", "vase", "state", "map"}
	anagram := convertToAnagramSlice(words)
	fmt.Println(anagram)
}

func convertToAnagramSlice(words []string) [][]string {
	anagram := [][]string{}
	visited := make(map[string]bool)
	for i := 0; i < len(words); i++ {
		if !visited[words[i]] {
			var newAnagram []string
			newAnagram = append(newAnagram, words[i])
			for j := i + 1; j < len(words); j++ {
				if checkAnagram(words[i], words[j]) {
					fmt.Println("anagram:", words[i], words[j])
					newAnagram = append(newAnagram, words[j])
					visited[words[j]] = true
				}
			}

			anagram = append(anagram, newAnagram)
		}

	}
	return anagram
}

func checkAnagram(anagram, word string) bool {
	s1 := strings.Split(anagram, "")
	s2 := strings.Split(word, "")
	sort(s1)
	sort(s2)
	anagram = strings.Join(s1, "")
	word = strings.Join(s2, "")
	return anagram == word
}

func sort(words []string) {
	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words); j++ {
			if words[i] < words[j] {
				words[i], words[j] = words[j], words[i]
			}
		}
	}
}
