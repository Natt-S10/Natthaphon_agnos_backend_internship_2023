package passwd

import (
	"fmt"
	"strings"
	"unicode"
	_ "unicode/utf8"
)

type RepeatRange struct {
	Begin int
	End   int
}

func (rR *RepeatRange) count() int {
	return rR.End - rR.Begin + 1
}

type PasswordObject struct {
	Original   string
	Tokenized  []string
	CountUpper int
	CountLower int
	CountDigit int
	CountSign  int
	RepeatList []RepeatRange
}

func (pO *PasswordObject) Init(password string) *PasswordObject {
	pO.Original = password
	pO.Tokenized = []string{}
	pO.CountUpper = 0
	pO.CountLower = 0
	pO.CountDigit = 0
	pO.CountSign = 0
	pO.RepeatList = []RepeatRange{}
	return pO
}

func (pO *PasswordObject) charTypeCount(c rune) {
	if unicode.IsLower(c) {
		pO.CountLower += 1
	} else if unicode.IsUpper(c) {
		pO.CountUpper += 1
	} else if unicode.IsDigit(c) {
		pO.CountDigit += 1
	} else if c == '!' || c == '.' {
		pO.CountSign += 1
	}
}

func (pO *PasswordObject) match(c rune, repeat RepeatRange) {
	dup := repeat.count()
	if dup >= repeatTreashold {
		newMatch := strings.Repeat(fmt.Sprintf("%c", c), dup)
		pO.Tokenized = append(pO.Tokenized, newMatch)
		pO.RepeatList = append(pO.RepeatList, repeat)
	} else {
		for i := 0; i < dup; i++ {
			pO.Tokenized = append(pO.Tokenized, fmt.Sprintf("%c", c))
		}
	}
}

func (pO *PasswordObject) tokenize() {
	var prev rune
	repeat := RepeatRange{}
	// append(s, encodeToUtf8(c))

	for idx, c := range pO.Original {
		// type count
		pO.charTypeCount(c)

		if idx == 0 {
			repeat.Begin = 0
			repeat.End = 0

		} else {
			if c == prev { // same val
				repeat.End = idx
			} else { // Match here
				// tokenize
				pO.match(prev, repeat)
				//reset
				repeat.Begin = idx
				repeat.End = idx
			}
		}
		prev = c
	}
	pO.match(prev, repeat)
}
