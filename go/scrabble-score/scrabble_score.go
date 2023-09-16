package scrabble

import (
	"strings"
)

// These are the letters.
var letterScores = map[int][]rune{
	1:  {'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T'},
	2:  {'D', 'G'},
	3:  {'B', 'C', 'M', 'P'},
	4:  {'F', 'H', 'V', 'W', 'Y'},
	5:  {'K'},
	8:  {'J', 'X'},
	10: {'Q', 'Z'},
}

func Score(word string) int {
	// create a more efficient representation of our letter scores
	var scoreMap = make(map[rune]int, 26)
	for s, letters := range letterScores {
		for _, l := range letters {
			scoreMap[l] = s
		}
	}

	word = strings.ToUpper(word)
	total := 0
	for _, c := range word {
		score, ok := scoreMap[c]
		if ok {
			total += score
		}
	}
	return total
}
