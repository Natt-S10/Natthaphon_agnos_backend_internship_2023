package passwd

// Criteria of Strong Password
const (
	// Length
	MINPASSWORDLEN = 6
	MAXPASSWORDLEN = 19

	// Char type : Presense of lowercase, uppercase, digit
	REQUIRELOWERCASE = true
	REQUIREUPPERCASE = true
	REQUIREDIGIT     = true
	REQUIRESIGN      = false
	// Assume total required types does not exceed MINPASSWORDLEN

	// repeating-free
	REPEATHRESHOLD = 3
)
