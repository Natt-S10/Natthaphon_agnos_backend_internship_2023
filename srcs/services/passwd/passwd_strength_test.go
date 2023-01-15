package passwd

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	testVal string
	label   bool
}

func TestPasswodLenghtChecking(t *testing.T) {
	// test cases
	lenList := []int{1, 2, 4, 5, 6, 7, 8, 10, 15, 18, 19, 20, 21, 22}
	labelList := []bool{false, false, false, false, true, true, true, true, true, true, true, false, false, false}

	for idx, i := range lenList {
		testpwd := strings.Repeat("a", i)
		result := IsAppropLength(testpwd)
		label := labelList[idx]
		assert.Equal(t, label, result, fmt.Sprintf("the password <<%s>> (len= %d) should be  evaluated as expected", testpwd, i))
		// if (res && !label) || (!res && label) {
		// 	t.Error("IsAppropLength of '", testpwd, "' should be", label, "but have", res)
		// }
	}
}

func TestPasswdCharTypeChecking(t *testing.T) {
	// test cases
	testSet := []testCase{
		{"h", false},
		{"H", false},
		{"5", false},
		{"hi", false},
		{"hi5", false},
		{"i5", false},
		{"Hi", false},
		{"Hi5", true},
		{"password", false},
		{"Password", false},
		{"passw0rd", false},
		{"Passw0rd", true},
	}

	for _, tC := range testSet {
		result := IsAppropCharType(tC.testVal)
		assert.Equal(t, tC.label, result, fmt.Sprintf("the password <<%s>> should be evaluated as expected", tC.testVal))
	}
}

func TestPasswdRepeatChecking(t *testing.T) {
	// test cases
	testSet := []testCase{
		{"abcd1234", true},
		{"abcc1234", true},
		{"Abcd1234", true},
		{"Abbb1234", false},
		{"Abcd1134", true},
		{"Abcd1114", false},
		{"aaAa1234", true},
		{"AAAd1234", false},
		{"11111111aabb", false},
	}

	for _, tC := range testSet {
		result := IsAppropRepeat(tC.testVal)
		assert.Equal(t, tC.label, result, fmt.Sprintf("the password <<%s>> should be evaluated as expected", tC.testVal))
	}
}
