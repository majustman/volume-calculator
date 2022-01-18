package figure

import (
	"fmt"
	"math"
)

type Ellipsoid struct {
	a, b, c float64
}

func (e *Ellipsoid) Volume() float64 {
	return (4 * math.Pi * e.a * e.b * e.c) / 3
}

func NewEllipsoid() *Ellipsoid {
	var e Ellipsoid
	fmt.Println("Please enter data:")
	for e.a == 0 || e.b == 0 || e.c == 0 {
		if e.a == 0 {
			fmt.Print("Radius A - ")
			fmt.Scan(&e.a)
		}
		if e.b == 0 {
			fmt.Print("Radius B - ")
			fmt.Scan(&e.b)
		}
		if e.c == 0 {
			fmt.Print("Radius C - ")
			fmt.Scan(&e.c)
		}
	}
	return &e
}
