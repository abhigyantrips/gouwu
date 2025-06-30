package gouwu

import (
	"regexp"
	"strings"
	"unicode"
)

// isAt checks if the value starts with '@' (for mentions/handles)
func isAt(value string) bool {
	if len(value) == 0 {
		return false
	}
	return value[0] == '@'
}

// getCapitalPercentage calculates what percentage of letters in the string are uppercase
func getCapitalPercentage(str string) float64 {
	var totalLetters, upperLetters int

	for _, r := range str {
		if !unicode.IsLetter(r) {
			continue
		}

		if unicode.IsUpper(r) {
			upperLetters++
		}
		totalLetters++
	}

	if totalLetters == 0 {
		return 0
	}

	return float64(upperLetters) / float64(totalLetters)
}

// isURI validates if the given string is a valid URI
// Direct port of the RFC 3986 validation logic from the original JS
func isURI(value string) bool {
	if value == "" {
		return false
	}

	// Check for illegal characters
	illegalChars := regexp.MustCompile(`[^a-zA-Z0-9:/?#\[\]@!$&'()*+,;=.\-_~%]`)
	if illegalChars.MatchString(value) {
		return false
	}

	// Check for incomplete hex escapes
	incompleteHex1 := regexp.MustCompile(`%[^0-9a-fA-F]`)
	incompleteHex2 := regexp.MustCompile(`%[0-9a-fA-F]([^0-9a-fA-F]|$)`)
	if incompleteHex1.MatchString(value) || incompleteHex2.MatchString(value) {
		return false
	}

	// RFC 3986 URI parsing regex - EXACTLY as in JS
	uriRegex := regexp.MustCompile(`(?:([^:/?#]+):)?(?://([^/?#]*))?([^?#]*)(?:\?([^#]*))?(?:#(.*))?`)
	matches := uriRegex.FindStringSubmatch(value)

	if matches == nil {
		return false
	}

	scheme := matches[1]
	authority := matches[2]
	path := matches[3]

	// Scheme and path are required, though the path can be empty
	// EXACT JS LOGIC: scheme && scheme.length && path.length >= 0
	if !(scheme != "" && len(path) >= 0) {
		return false
	}

	// If authority is present, path must be empty or start with /
	if authority != "" {
		if !(len(path) == 0 || strings.HasPrefix(path, "/")) {
			return false
		}
	} else {
		// If no authority, path must not start with //
		if strings.HasPrefix(path, "//") {
			return false
		}
	}

	// Scheme validation: must start with letter, then letters/digits/+/./-
	schemeRegex := regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9+\-.]*$`)
	if !schemeRegex.MatchString(scheme) {
		return false
	}

	return true
}

// isBreak checks if the word is just whitespace
func isBreak(word string) bool {
	return strings.TrimSpace(word) == ""
}
