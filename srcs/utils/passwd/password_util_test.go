package passwd

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

const testPassword = "PassworrrrdSSS12!!"

func TestPasswordInit(t *testing.T) {
	newPassswordObject := PasswordObject{}
	newPassswordObject.Init(testPassword)

	expected := PasswordObject{
		Original:   "PassworrrrdSSS12!!",
		Tokenized:  []string{},
		CountUpper: 0,
		CountLower: 0,
		CountDigit: 0,
		CountSign:  0,
		RepeatList: []RepeatRange{},
	}
	assert.Equal(t, expected, newPassswordObject)
}

type testCase struct {
	testInput rune
	expectedL int
	expectedU int
	expectedD int
	expectedS int
}

func TestCharTypeCount(t *testing.T) {
	newPassswordObject := PasswordObject{}
	newPassswordObject.Init(testPassword)

	testSet := []testCase{
		{'#', 0, 0, 0, 0},
		{'a', 1, 0, 0, 0},
		{'a', 2, 0, 0, 0},
		{'A', 2, 1, 0, 0},
		{'A', 2, 2, 0, 0},
		{'3', 2, 2, 1, 0},
		{'g', 3, 2, 1, 0},
		{'!', 3, 2, 1, 1},
		{'!', 3, 2, 1, 2},
		{'.', 3, 2, 1, 3},
		{'1', 3, 2, 2, 3},
		{'9', 3, 2, 3, 3},
	}
	for idx, tC := range testSet {
		newPassswordObject.charTypeCount(tC.testInput)
		assert.Equal(t, tC.expectedL, newPassswordObject.CountLower, fmt.Sprintf("checkLower round %d", idx))
		assert.Equal(t, tC.expectedU, newPassswordObject.CountUpper, fmt.Sprintf("checkUpper round %d", idx))
		assert.Equal(t, tC.expectedD, newPassswordObject.CountDigit, fmt.Sprintf("checkDigit round %d", idx))
		assert.Equal(t, tC.expectedS, newPassswordObject.CountSign, fmt.Sprintf("checkSign round %d", idx))
	}
}

func TestPasswordObjectMatch(t *testing.T) {
	newPassswordObject := PasswordObject{}
	newPassswordObject.Init(testPassword)
	newPassswordObject.match('g', RepeatRange{2, 3})
	assert.Equal(t, []RepeatRange{}, newPassswordObject.RepeatList, "test RepeatList 1")
	assert.Equal(t, []string{"g", "g"}, newPassswordObject.Tokenized, "test Tokenized 1")
	newPassswordObject.match('P', RepeatRange{2, 2})
	assert.Equal(t, []RepeatRange{}, newPassswordObject.RepeatList, "test RepeatList 2")
	assert.Equal(t, []string{"g", "g", "P"}, newPassswordObject.Tokenized, "test Tokenized 2")
	newPassswordObject.match('x', RepeatRange{2, 4})
	assert.Equal(t, []RepeatRange{{2, 4}}, newPassswordObject.RepeatList, "test RepeatList 3")
	assert.Equal(t, []string{"g", "g", "P", "xxx"}, newPassswordObject.Tokenized, "test Tokenized  3")
	newPassswordObject.match('!', RepeatRange{5, 6})
	assert.Equal(t, []RepeatRange{{2, 4}}, newPassswordObject.RepeatList, "test RepeatList 4")
	assert.Equal(t, []string{"g", "g", "P", "xxx", "!", "!"}, newPassswordObject.Tokenized, "test Tokenized  4")
}

func TestPasswordObjectTokenize(t *testing.T) {
	newPassswordObject := PasswordObject{}
	newPassswordObject.Init(testPassword)
	newPassswordObject.tokenize()

	expected := PasswordObject{
		Original:   "PassworrrrdSSS12!!",
		Tokenized:  []string{"P", "a", "s", "s", "w", "o", "rrrr", "d", "SSS", "1", "2", "!", "!"},
		CountUpper: 4,
		CountLower: 10,
		CountDigit: 2,
		CountSign:  2,
		RepeatList: []RepeatRange{{6, 9}, {11, 13}},
	}
	assert.Equal(t, expected, newPassswordObject)
}
