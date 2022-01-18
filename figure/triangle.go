package figure

import (
	"math"
)

type Triangle struct {
	SideA float64
	SideB float64
	SideC float64
}

func (t *Triangle) Area() float64 {
	s := (t.SideA + t.SideB + t.SideC) / 2
	return math.Sqrt(s * (s - t.SideA) * (s - t.SideB) * (s - t.SideC))
}
