package figure

import (
	"fmt"
	"math"
)

type Sphere struct {
	Radius float64
}

func (s *Sphere) Volume() float64 {
	return (4 * math.Pi * s.Radius * s.Radius * s.Radius) / 3
}

func NewSphere() *Sphere {
	var s Sphere
	fmt.Println("Please enter data:")
	for s.Radius == 0 {
		if s.Radius == 0 {
			fmt.Print("Radius - ")
			fmt.Scan(&s.Radius)
		}
	}
	return &s
}
