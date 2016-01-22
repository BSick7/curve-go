package types

type FillRule int8
const (
	FillRuleEvenOdd = iota
	FillRuleNonZero = iota
)

type SweepDirection int8
const (
	Counterclockwise = iota
	Clockwise = iota
)

type ISegmentExecutor interface {
	Exec(runner ISegmentRunner)
}

type ISegmentRunner interface {
	SetFillRule(fill_rule FillRule)
	ClosePath()
	MoveTo(x float32, y float32)
	LineTo(x float32, y float32)
	BezierCurveTo(cp1x float32, cp1y float32, cp2x float32, cp2y float32, x float32, y float32)
	QuadraticCurveTo(cpx float32, cpy float32, x float32, y float32)
	Arc(x float32, y float32, radius float32, startAngle float32, endAngle float32, anticlockwise bool)
	ArcTo(x1 float32, y1 float32, x2 float32, y2 float32, radius float32)
	Ellipse(cx float32, cy float32, rx float32, ry float32, rotation float32, startAngle float32, endAngle float32, antiClockwise bool)
}

type ISegment func(runner ISegmentRunner)
