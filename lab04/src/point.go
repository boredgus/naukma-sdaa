package src

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

func NewPoint(x, y float64) Point {
	return Point{x, y}
}

func (p Point) String() string {
	return fmt.Sprintf("(%.2f, %.2f)", p.X, p.Y)
}

func (p Point) Transplant(max, width float64) (float64, float64) {
	return p.transplantValue(p.X, max, width, false), p.transplantValue(p.Y, max, width, true)
}

func (p Point) transplantValue(value, max, width float64, isReversed bool) (res float64) {
	halfWidth := width / 2

	if isReversed {
		if value < 0 {
			res = (1 + (math.Abs(value) / max)) * halfWidth
		} else {
			res = (max - value) / max * halfWidth
		}
	} else {
		if value < 0 {
			res = (max + value) / max * halfWidth
		} else {
			res = (1 + (value / max)) * halfWidth
		}
	}

	return
}
