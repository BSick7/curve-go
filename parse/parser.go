package parse

import (
	"fmt"
	"io"
	"strconv"
	"github.com/wsick/curve-go/types"
)

type Parser struct {
	s    *Scanner
	pos  int
	curx float64
	cury float64
	buf  struct {
			 tok Token
			 lit string
			 n   int
		 }
}

type point struct {
	X float64
	Y float64
}

func NewParser(r io.Reader) *Parser {
	return &Parser{
		s: NewScanner(r),
		pos: 0,
		curx: 0,
		cury: 0,
	}
}

func (p *Parser) Parse(runner types.ISegmentRunner) error {
	for {
		if ok, err := p.scanCommand(runner); err != nil {
			return fmt.Errorf("[pos: %d]: %s", p.pos - len(p.buf.lit), err)
		} else if !ok {
			return nil
		}
	}
}

// scan returns the next token from the underlying scanner.
// If a token has been unscanned then read that instead.
func (p *Parser) scan() (tok Token, lit string) {
	// If we have a token on the buffer, then return it.
	if p.buf.n != 0 {
		p.buf.n = 0
		p.pos += len(p.buf.lit)
		return p.buf.tok, p.buf.lit
	}

	// Otherwise read the next token from the scanner.
	tok, lit = p.s.Scan()

	// Save it to the buffer in case we unscan later.
	p.buf.tok, p.buf.lit = tok, lit
	p.pos += len(p.buf.lit)

	return
}

// unscan pushes the previously read token back onto the buffer.
func (p *Parser) unscan() {
	p.buf.n = 1
	p.pos -= len(p.buf.lit)
}

// scanIgnoreWhitespace scans the next non-whitespace token.
func (p *Parser) scanIgnoreWhitespace() (tok Token, lit string) {
	tok, lit = p.scan()
	if tok == WS {
		tok, lit = p.scan()
	}
	return
}

func (p *Parser) scanFlag() (int8, error) {
	tok, lit := p.scanIgnoreWhitespace()
	if tok != DIGITS {
		return 0, fmt.Errorf("found %q, expected flag (0 or 1)", lit)
	}

	if i, err := strconv.ParseInt(lit, 10, 8); err != nil {
		return 0, fmt.Errorf("found %q, expected flag (0 or 1)", lit)
	} else {
		return int8(i), nil
	}
}

func (p *Parser) scanPoint() (*point, error) {
	x, err := p.scanNumber()
	if err != nil {
		return nil, err
	}

	tok, lit := p.scanIgnoreWhitespace()
	if tok != COMMA && tok != WS {
		return nil, fmt.Errorf("found %q, expected whitespace between points", lit)
	}

	y, err := p.scanNumber()
	if err != nil {
		return nil, err
	}

	return &point{
		X: x,
		Y: y,
	}, nil
}
