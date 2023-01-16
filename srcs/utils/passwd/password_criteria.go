package passwd

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
