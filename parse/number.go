package parse
import (
	"bytes"
	"math"
	"fmt"
	"strconv"
)

func (p *Parser) scanNumber() (float64, error) {
	var buf bytes.Buffer

	// Place the next non-ws token on the parser buffer for scanOptionalSign
	p.scanIgnoreWhitespace()
	p.unscan()

	// Parse leading +/- sign
	sign := p.scanOptionalSign()

	// Try Infinity, -Infinity, NaN
	if s, ok := p.scanSpecialNumber(sign); ok {
		return s, nil
	}

	// Parse whole digits
	if tok, lit := p.scan(); tok == DIGITS {
		buf.WriteString(lit)
	} else {
		return 0, fmt.Errorf("found %q, expected number", lit)
	}

	// Parse fraction digits
	if tok, lit := p.scan(); tok == PERIOD {
		buf.WriteString(lit)
		if tok2, lit2 := p.scan(); tok2 == DIGITS {
			buf.WriteString(lit2)
		} else {
			return 0, fmt.Errorf("found %q, expected fraction digits", lit2)
		}
	} else {
		p.unscan()
	}

	if tok, lit := p.scan(); tok == EXPONENT {
		buf.WriteString(lit)
		if tok2, lit2 := p.scan(); tok2 == DIGITS {
			buf.WriteString(lit2)
		} else {
			return 0, fmt.Errorf("found %q, expected exponential digits", lit2)
		}
	}

	if u, err := strconv.ParseFloat(buf.String(), 32); err != nil {
		return 0, err
	} else {
		return u, nil
	}
}

func (p *Parser) scanOptionalSign() int {
	if tok, _ := p.scan(); tok == NEGATIVE {
		return -1
	} else if tok == POSITIVE {
		return 1
	} else {
		p.unscan()
	}
	return 0
}

func (p *Parser) scanSpecialNumber(sign int) (float64, bool) {
	if tok, lit := p.scan(); tok == LETTERS {
		if lit == "Infinity" {
			return math.Inf(sign), true
		} else if lit == "NaN" {
			return math.NaN(), true
		} else {
			p.unscan()
		}
	}
	return 0, false
}

func degreesToRadians(angle float64) float64 {
	return angle / 180.0 * math.Pi;
}
