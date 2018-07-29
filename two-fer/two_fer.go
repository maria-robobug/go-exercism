package twofer

import (
	"strings"
	"unicode"
)

var message string = "One for %s, one for me."

// ShareWith - returns "One for X, one for me" message
func ShareWith(name string) string {
	s := strings.TrimFunc(name, unicode.IsSpace)

	if len(s) == 0 {
		return strings.Replace(message, "%s", "you", 1)
	}

	return strings.Replace(message, "%s", name, 1)
}
