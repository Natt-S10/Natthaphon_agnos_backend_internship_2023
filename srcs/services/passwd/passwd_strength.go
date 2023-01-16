package passwd

import (
	_ "fmt"
	"unicode"

	utils "github.com/Natt-S10/Natthaphon_agnos_backend_internship_2023/srcs/utils/passwd"
)

type CharTypePresence struct {
	lowercase bool
	uppercase bool
	digit     bool
}

type RepeatPresense struct {
	repeatBegin int
	repeatEnd   int
}

func (cp *CharTypePresence) isComplete() bool {
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

// func GetRepeatPresence(password string) RepeatPresense {
// 	// return
// 	return RepeatPresense{}
// }

func IsAppropLength(password string) bool {
	return utils.MINPASSWORDLEN <= len(password) && len(password) <= utils.MAXPASSWORDLEN
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
		if repeat >= utils.REPEATHRESHOLD {
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
