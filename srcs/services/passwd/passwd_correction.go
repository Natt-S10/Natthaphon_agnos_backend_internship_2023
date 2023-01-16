package passwd

import (
	utils "github.com/Natt-S10/Natthaphon_agnos_backend_internship_2023/srcs/utils/passwd"
)

// func shrinkRepeat(dup int) (int,int) { //(steps, reduced alphabet count)
// 	if utils.maxPasswordLen - dup +
// }

func totalMissing(pO *utils.PasswordObject, passwdLen int) int {
	typeMissing := 0
	if pO.IsMissingLower() {
		typeMissing++
	}
	if pO.IsMissingUpper() {
		typeMissing++
	}
	if pO.IsMissingDigit() {
		typeMissing++
	}
	if pO.IsMissingSign() {
		typeMissing++
	}

	charMissing := utils.MINPASSWORDLEN - passwdLen

	if typeMissing > charMissing {
		return typeMissing
	}
	return charMissing
}

func PasswordCorrectSteps(password string) int {
	passwordObject := utils.PasswordObject{}
	passwordObject.Init(password)

	passwdLen := len(password)
	steps := 0

	repeatLeft := []int{}
	for _, repeat := range passwordObject.RepeatList { // shrink to fit the max limit
		if passwdLen > utils.MAXPASSWORDLEN { // too long
			var shrinked int
			if utils.MAXPASSWORDLEN < passwdLen-repeat.Count()+utils.REPEATHRESHOLD-1 { // shrink as much as it could, still to long
				shrinked = repeat.Count() - utils.REPEATHRESHOLD + 1
			} else { // shrink just enough
				shrinked = utils.MAXPASSWORDLEN - passwdLen
				leftOvered := repeat.Count() - shrinked
				if leftOvered >= utils.REPEATHRESHOLD { // still Repeat, even shrinked
					repeatLeft = append(repeatLeft, leftOvered)
				}
			}
			steps += shrinked
			passwdLen -= shrinked
		}
	}

	for _, lefted := range repeatLeft {
		replaceable := lefted / utils.REPEATHRESHOLD // if 3 or more is repeat, the every 3 letter can be replaced to avoid repeat. ex. -> aaXaaX
		passwordObject.ReplaceN(replaceable)
		steps += replaceable
	}

	steps += totalMissing(&passwordObject, passwdLen)
	return steps
}
