package parse

// Token represents a lexical token.
type Token int

const (
	// Special tokens
	ILLEGAL Token = iota
	EOF
	WS

	// Letter
	LETTERS

	// Digit
	EXPONENT // E, e
	DIGITS
	INFINITY // Infinity, -Infinity
	NAN		 // NaN

	// Misc characters
	PERIOD	 // .
	COMMA    // ,
	QUOTE	 // '
	POSITIVE // +
	NEGATIVE // -
)
