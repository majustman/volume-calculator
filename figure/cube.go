package figure

import "fmt"

type Cube struct {
	Height float64
}

func (c *Cube) Volume() float64 {
	return c.Height * c.Height * c.Height
}

func NewCube() *Cube {
	var c Cube
	fmt.Println("Please enter data:")
	for c.Height == 0 {
		if c.Height == 0 {
			fmt.Print("Height - ")
			fmt.Scan(&c.Height)
		}
	}
	return &c
}
