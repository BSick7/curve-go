package parse

import (
	"github.com/wsick/curve-go/types"
	"fmt"
)

func (p *Parser) scanCommand(runner types.ISegmentRunner) (bool, error) {
	tok, lit := p.scanIgnoreWhitespace()
	if tok == EOF {
		return false, nil
	}
	if tok != LETTERS {
		return true, fmt.Errorf("found %q, expected command", lit)
	}

	switch lit {
	case "f":
	case "F":
		return true, p.scanFillRule(runner)
	case "m":
		return true, p.scanMoveTo(runner, true)
	case "M":
		return true, p.scanMoveTo(runner, false)
	case "l":
		return true, p.scanLineTo(runner, true)
	case "L":
		return true, p.scanLineTo(runner, false)
	case "h":
		return true, p.scanHorizontalLineTo(runner, true)
	case "H":
		return true, p.scanHorizontalLineTo(runner, false)
	case "v":
		return true, p.scanVerticalLineTo(runner, true)
	case "V":
		return true, p.scanVerticalLineTo(runner, false)
	case "c":
		return true, p.scanCubicBezierTo(runner, true)
	case "C":
		return true, p.scanCubicBezierTo(runner, false)
	case "s":
		return true, p.scanSmoothCubicBezierTo(runner, true)
	case "S":
		return true, p.scanSmoothCubicBezierTo(runner, false)
	case "q":
		return true, p.scanQuadraticBezierTo(runner, true)
	case "Q":
		return true, p.scanQuadraticBezierTo(runner, false)
	case "t":
		return true, p.scanSmoothQuadraticBezierTo(runner, true)
	case "T":
		return true, p.scanSmoothQuadraticBezierTo(runner, false)
	case "z":
	case "Z":
		return true, p.scanClosePath(runner)
	}

	return true, nil
}

func (p *Parser) scanFillRule(runner types.ISegmentRunner) error {
	_, lit := p.scanIgnoreWhitespace()
	switch lit {
	case "0":
		runner.SetFillRule(types.FillRuleEvenOdd)
	case "1":
		runner.SetFillRule(types.FillRuleNonZero)
	}
	return fmt.Errorf("found %q, expected fill rule of 0 or 1", lit)
}

func (p *Parser) scanMoveTo(runner types.ISegmentRunner, rel bool) error {
	p1, err := p.scanPoint()
	if err != nil {
		return err
	}
	if rel {
		p1.X += p.curx
		p1.Y += p.cury
	}
	runner.MoveTo(p1.X, p1.Y)
	p.curx = p1.X
	p.cury = p1.Y
	//TODO: Handle multiple points
	return nil
}

func (p *Parser) scanLineTo(runner types.ISegmentRunner, rel bool) error {
	p1, err := p.scanPoint()
	if err != nil {
		return err
	}
	if rel {
		p1.X += p.curx
		p1.Y += p.cury
	}
	runner.LineTo(p1.X, p1.Y)
	p.curx = p1.X
	p.cury = p1.Y
	//TODO: Handle multiple points
	return nil
}

func (p *Parser) scanHorizontalLineTo(runner types.ISegmentRunner, rel bool) error {
	x, err := p.scanNumber()
	if err != nil {
		return err
	}
	if rel {
		x += p.curx
	}
	runner.LineTo(x, p.cury)
	p.curx = x
	return nil
}

func (p *Parser) scanVerticalLineTo(runner types.ISegmentRunner, rel bool) error {
	y, err := p.scanNumber()
	if err != nil {
		return err
	}
	if rel {
		y += p.curx
	}
	runner.LineTo(p.curx, y)
	p.cury = y
	return nil
}

func (p *Parser) scanCubicBezierTo(runner types.ISegmentRunner, rel bool) error {
	cp1, err := p.scanPoint()
	if err != nil {
		return err
	}
	if rel {
		cp1.X += p.curx
		cp1.Y += p.cury
	}
	cp2, err := p.scanPoint()
	if err != nil {
		return err
	}
	if rel {
		cp2.X += p.curx
		cp2.Y += p.cury
	}
	p1, err := p.scanPoint()
	if err != nil {
		return err
	}
	if rel {
		p1.X += p.curx
		p1.Y += p.cury
	}
	runner.BezierCurveTo(cp1.X, cp1.Y, cp2.X, cp2.Y, p1.X, p1.Y)
	p.curx = p1.X
	p.cury = p1.Y
	//TODO: Handle multiple points
	return nil
}

func (p *Parser) scanSmoothCubicBezierTo(runner types.ISegmentRunner, rel bool) error {
	//TODO: Implement
	return nil
}

func (p *Parser) scanQuadraticBezierTo(runner types.ISegmentRunner, rel bool) error {
	cp, err := p.scanPoint()
	if err != nil {
		return err
	}
	if rel {
		cp.X += p.curx
		cp.Y += p.cury
	}
	p1, err := p.scanPoint()
	if err != nil {
		return err
	}
	if rel {
		p1.X += p.curx
		p1.Y += p.cury
	}
	runner.QuadraticCurveTo(cp.X, cp.Y, p1.X, p1.Y)
	p.curx = p1.X
	p.cury = p1.Y
	return nil
}

func (p *Parser) scanSmoothQuadraticBezierTo(runner types.ISegmentRunner, rel bool) error {
	//TODO: Implement
	return nil
}

func (p *Parser) scanEllipticalArcTo(runner types.ISegmentRunner, rel bool) error {
	r, err := p.scanPoint()
	if err != nil {
		return err
	}
	angle, err := p.scanNumber()
	phi := degreesToRadians(angle)
	if err != nil {
		return err
	}
	fa, err := p.scanFlag()
	if err != nil {
		return err
	}
	fs, err := p.scanFlag()
	if err != nil {
		return err
	}
	p1, err := p.scanPoint()
	if err != nil {
		return err
	}
	if rel {
		p1.X += p.curx
		p1.Y += p.cury
	}

	runner.EllipticalArc(r.X, r.Y, phi, fa, fs, p1.X, p1.Y)
	p.curx = p1.X
	p.cury = p1.Y
	return nil
}

func (p *Parser) scanClosePath(runner types.ISegmentRunner) error {
	runner.ClosePath()
	return nil
}
