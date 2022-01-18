package figure

import (
	"fmt"
	"math"
)

type Cylinder struct {
	Radius float64
	Height float64
}

func (c *Cylinder) Volume() float64 {
	return c.Height * math.Pi * c.Radius * c.Radius
}

func NewCylinder() *Cylinder {
	var c Cylinder
	fmt.Println("Please enter data:")
	for c.Radius == 0 || c.Height == 0 {
		if c.Radius == 0 {
			fmt.Print("Radius - ")
			fmt.Scan(&c.Radius)
		}
		if c.Height == 0 {
			fmt.Print("Height - ")
			fmt.Scan(&c.Height)
		}
	}
	return &c
}
