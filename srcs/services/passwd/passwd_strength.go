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

type CharTypePresence struct {
	lowercase 	bool
	uppercase 	bool
	digit 		bool
}

type RepeatPresense struct {
	repeatBegin	int
	repeatEnd	int
}

func (cp *CharTypePresence) isComplete () bool {
	return cp.lowercase && cp.uppercase && cp.digit
}

func GetCharTypePresence(password string) CharTypePresence {
	contains := CharTypePresence{false, false, false}
	for _, c := range password {
		if unicode.IsLower(c) {
			contains.lowercase = true
		}
		if unicode.IsUpper(c) {
			contains.uppercase = true
		}
		if unicode.IsDigit(c) {
			contains.digit = true
		}

		// shortcutting
		if contains.isComplete() {
			break
		}
	}
	return contains
}

func GetRepeatPresence(password string) RepeatPresense {
	// 
}

func IsAppropLength(password string) bool {
	return minPasswordLen <= len(password) && len(password) <= maxPasswordLen
}

func IsAppropCharType(password string) bool {
	presense := GetCharTypePresence(password)
	return presense.isComplete()
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


// var prev rune
// repeat := 1
// for idx, c := range password {
// 	if idx != 0 {
// 		if c == prev {
// 			repeat += 1
// 			rP.repeatEnd = idx
// 		} else {
// 			repeat = 1
// 			rP.repeatBegin = idx+1
// 		}
// 	}
// 	if repeat >= repeatTreashold {
// 		rP.isRepeat = true
// 		break
// 	}
// 	prev = c
// }
// return rP