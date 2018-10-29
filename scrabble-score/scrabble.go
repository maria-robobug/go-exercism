package scrabble

import (
	"strings"
)

// Score - Given a word, compute scrabble scrore
func Score(word string) (score int) {
	table := map[int][]string{
		1:  []string{"A", "E", "I", "O", "U", "L", "N", "R", "S", "T"},
		2:  []string{"D", "G"},
		3:  []string{"B", "C", "M", "P"},
		4:  []string{"F", "H", "V", "W", "Y"},
		5:  []string{"K"},
		8:  []string{"J", "X"},
		10: []string{"Q", "Z"},
	}

	for _, c := range word {
		for k, v := range table {
			for _, l := range v {
				if strings.ToLower(string(c)) == strings.ToLower(l) {
					score += k
				}
			}
		}
	}

	return score
}
