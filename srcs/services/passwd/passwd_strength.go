package passwd

import (
	_ "fmt"
	"unicode"
)

// Criteria of Strong Password
const (
	// Length
	minPasswordLen = 6
	maxPasswordLen = 19

	// Char type : Presense of lowercase, uppercase, digit
	requireLowercase = true
	requireUppercase = true
	requireDigit     = true

	// repeating-free
	repeatTreashold = 3
)

func IsAppropLength(password string) bool {
	return minPasswordLen <= len(password) && len(password) <= maxPasswordLen
}

func IsAppropCharType(password string) bool {
	containsLower := false
	containsUpper := false
	containsDigit := false

	for _, c := range password {
		if unicode.IsLower(c) {
			containsLower = true
		}
		if unicode.IsUpper(c) {
			containsUpper = true
		}
		if unicode.IsDigit(c) {
			containsDigit = true
		}

		// shortcutting
		if containsLower && containsUpper && containsDigit {
			return true
		}
	}
	return false
}

func IsAppropRepeat(password string) bool {
	var prev rune
	repeat := 1
	for idx, c := range password {
		if idx != 0 {
			if c == prev {
				repeat += 1
			} else {
				repeat = 1
			}
		}
		if repeat >= repeatTreashold {
			return false
		}
		prev = c
	}
	return true
}
