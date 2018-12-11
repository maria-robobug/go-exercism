package isogram

import (
	"unicode"
)

// IsIsogram - determines if word is an isogram
func IsIsogram(input string) bool {
	var count int
	seen := map[rune]bool{}

	for _, c := range input {
		if !unicode.IsLetter(c) {
			continue
		}

		c = unicode.ToLower(c)
		if seen[c] {
			count++
		}
		seen[c] = true
	}

	return count == 0
}
