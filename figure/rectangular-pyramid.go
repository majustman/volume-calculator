package figure

import "fmt"

type RectangularPyramid struct {
	*Rectangle
	Height float64
}

func (rp *RectangularPyramid) Volume() float64 {
	return (rp.Rectangle.Area() * rp.Height) / 3
}

func NewRectangularPyramid() *RectangularPyramid {
	var rp RectangularPyramid
	rp.Rectangle = &Rectangle{}
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
