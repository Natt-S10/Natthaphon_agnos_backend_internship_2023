package passwd

import (
	"fmt"
	"testing"

	utils "github.com/Natt-S10/Natthaphon_agnos_backend_internship_2023/srcs/utils/passwd"
	"github.com/stretchr/testify/assert"
)

func TestTotalMissing(t *testing.T) {
	newPO := utils.PasswordObject{}
	type tuple struct {
		Str    string
		expect int
	}
	cases := []tuple{
		{"Pa1", 3},
		{"pa.1", 2},
		{"Password", 1},
		{"p", 5},
		{"passwo", 2},
		{"passworD0", 0},
	}
	for _, cs := range cases {
		newPO.Init(cs.Str).TokenizePassword()
		assert.Equal(t, cs.expect, totalMissing(&newPO, len(cs.Str)), fmt.Sprintf("testing password: %s", cs.Str))
	}
}

func TestPasswordCorrectSteps(t *testing.T) {
	type pair struct {
		input  string
		expect int
	}
	cases := []pair{
		{"Passw0rd", 0},
		{"passw0rd", 1},
		{"password", 2},
		{"Passw0rd", 0},
		{"0123456789012345aAA", 0},
		{"0123456789012345aAAB", 1},
		{"0123456789012345aAAA", 1},
		{"0123456789012345aAAAA", 2},
		{"0123456789012345aAAAAB", 3},
		{"0123456789012345AAABB", 3},
		{"AAABBBCCC111222...", 6},     // AAgBBgCCg11g22g..g 6
		{"AAABBBCCC111222...xxx", 7},  // AAgBBgCCg11g22g..gx 8
		{"1234567890123456789012", 5}, //remove 3 then replace 2
	}
	for _, cs := range cases {
		assert.Equal(t, cs.expect, PasswordCorrectSteps(cs.input), cs.input)
	}
}
