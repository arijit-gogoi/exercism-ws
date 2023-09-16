package isogram

import (
	"strings"
	"unicode"
)

// func IsIsogram(word string) bool {
// 	m := make(map[rune]int, len(word))
// 	word = strings.ToLower(word)
// 	runes := []rune(word)
// 	for _, r := range runes {
// 		if r != ' ' && r != '-' {
// 			m[r]++
// 		}

// 		if m[r] > 1 {
// 			return false
// 		}
// 	}
// 	return true
// }

func IsIsogram(s string) bool {
	s = strings.ToLower(s)
	for i, c := range s {
		if unicode.IsLetter(c) && strings.ContainsRune(s[i+1:], c) {
			return false
		}
	}
	return true
}
