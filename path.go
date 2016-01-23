package curve

import (
//	"log"
	"bytes"
	"github.com/wsick/curve-go/parse"
	"github.com/wsick/curve-go/types"
)

type Path struct{}

func NewPath(data string) (*Path, error) {
	var buf bytes.Buffer
	buf.WriteString(data)
	parser := parse.NewParser(bytes.NewReader(buf.Bytes()))
	path := &Path{}
	if err := parser.Parse(path); err != nil {
		return nil, err
	}
	return path, nil
}

func (p *Path) SetFillRule(fill_rule types.FillRule) {
	//log.Println("setFillRule");
}
func (p *Path) ClosePath() {
	//log.Println("closePath");
}
func (p *Path) MoveTo(x float64, y float64) {
	//log.Println("moveTo");
}
func (p *Path) LineTo(x float64, y float64) {
	//log.Println("lineTo");
}
func (p *Path) BezierCurveTo(cp1x float64, cp1y float64, cp2x float64, cp2y float64, x float64, y float64) {
	//log.Println("bezierCurveTo");
}
func (p *Path) QuadraticCurveTo(cpx float64, cpy float64, x float64, y float64) {
	//log.Println("quadraticCurveTo");
}
func (p *Path) EllipticalArc(rx float64, ry float64, phi float64, fa int8, fs int8, x float64, y float64) {
	//log.Println("ellipticalArc");
}
