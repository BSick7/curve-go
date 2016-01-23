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
	MoveTo(x float64, y float64)
	LineTo(x float64, y float64)
	BezierCurveTo(cp1x float64, cp1y float64, cp2x float64, cp2y float64, x float64, y float64)
	QuadraticCurveTo(cpx float64, cpy float64, x float64, y float64)
	//Arc(x float64, y float64, radius float64, startAngle float64, endAngle float64, anticlockwise bool)
	//ArcTo(x1 float64, y1 float64, x2 float64, y2 float64, radius float64)
	EllipticalArc(rx float64, ry float64, phi float64, fa int8, fs int8, x float64, y float64)
}

type ISegment func(runner ISegmentRunner)
