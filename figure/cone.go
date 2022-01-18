package figure

import (
	"fmt"
	"math"
)

type Cone struct {
	Radius float64
	Height float64
}

func (c *Cone) Volume() float64 {
	return (math.Pi * c.Radius * c.Radius * c.Height) / 3
}

func NewCone() *Cone {
	var c Cone
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
