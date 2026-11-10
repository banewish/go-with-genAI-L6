// Package validate provides simple syntactic validators for common
// user-input formats.
//
// Homework — Task 2 (Lesson 6: File I/O, JSON and Testing):
// Implement ValidateEmail and/or ValidatePhone below (your mentor may
// ask for just one) and extend the test tables in validate_test.go to
// at least 8 cases each, including edge cases.
package validate

import (
	"strings"
	"unicode"
)

// ValidateEmail reports whether s is a syntactically valid email address.
//
// Accepts: local@domain format where local and domain are non-empty.
// Rejects: empty strings, multiple @, whitespace, consecutive dots, trailing dot.
func ValidateEmail(s string) bool {
	if s == "" || strings.ContainsAny(s, " \t\n\r") {
		return false
	}

	parts := strings.Split(s, "@")
	if len(parts) != 2 {
		return false
	}

	local, domain := parts[0], parts[1]

	if local == "" || domain == "" {
		return false
	}

	if strings.Contains(local, "..") || strings.Contains(domain, "..") {
		return false
	}

	if strings.HasSuffix(domain, ".") {
		return false
	}

	if !strings.Contains(domain, ".") {
		return false
	}

	return true
}

// ValidatePhone reports whether s is a syntactically valid phone number.
//
// Accepts: +<digits>, <digits> with optional - separators, at least 7 digits total.
// Rejects: empty strings, letters, spaces, less than 7 digits.
func ValidatePhone(s string) bool {
	if s == "" || strings.ContainsAny(s, " \t\n\r") {
		return false
	}

	for _, ch := range s {
		if !unicode.IsDigit(ch) && ch != '+' && ch != '-' {
			return false
		}
	}

	digitCount := 0
	for _, ch := range s {
		if unicode.IsDigit(ch) {
			digitCount++
		}
	}

	return digitCount >= 7
}
