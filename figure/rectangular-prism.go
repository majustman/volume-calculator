package figure

import "fmt"

type RectangularPrism struct {
	Length float64
	Width  float64
	Height float64
}

func (r *RectangularPrism) Volume() float64 {
	return r.Width * r.Length * r.Height
}

func NewRectangularPrism() *RectangularPrism {
	var rp RectangularPrism
	fmt.Println("Please enter data:")
	for rp.Height == 0 || rp.Length == 0 || rp.Width == 0 {
		if rp.Length == 0 {
			fmt.Print("Length - ")
			fmt.Scan(&rp.Length)
		}
		if rp.Width == 0 {
			fmt.Print("Width - ")
			fmt.Scan(&rp.Width)
		}
		if rp.Height == 0 {
			fmt.Print("Height - ")
			fmt.Scan(&rp.Height)
		}
	}
	return &rp
}
