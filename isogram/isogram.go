package isogram

import (
	"unicode"
)

// IsIsogram - determines if word is an isogram
func IsIsogram(input string) bool {
	var count int
	seen := make(map[string]int)

	for _, c := range input {
		if !unicode.IsLetter(c) {
			continue
		}

		c = unicode.ToLower(c)
		char := string(c)
		if _, ok := seen[char]; ok == true {
			count++
		}
		seen[char] = count
	}

	return count == 0
}
