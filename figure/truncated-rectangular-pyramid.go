package figure

import "fmt"

type TruncatedRectangularPyramid struct {
	TopRectangle  *Rectangle
	BaseRectangle *Rectangle
	Height        float64
}

func (trp *TruncatedRectangularPyramid) Volume() float64 {
	a := trp.TopRectangle.Width
	b := trp.TopRectangle.Length
	A := trp.BaseRectangle.Width
	B := trp.BaseRectangle.Length
	h := trp.Height
	return ((A*b + a*B + 2*(a*b+A*B)) * h) / 6
}

func NewTruncatedRectangularPyramid() *TruncatedRectangularPyramid {
	var trp TruncatedRectangularPyramid
	trp.TopRectangle = &Rectangle{}
	trp.BaseRectangle = &Rectangle{}
	fmt.Println("Please enter data of Top Rectangle:")
	for trp.TopRectangle.Length == 0 || trp.TopRectangle.Width == 0 {
		if trp.TopRectangle.Length == 0 {
			fmt.Print("Length - ")
			fmt.Scan(&trp.TopRectangle.Length)
		}
		if trp.TopRectangle.Width == 0 {
			fmt.Print("Width - ")
			fmt.Scan(&trp.TopRectangle.Width)
		}
	}
	fmt.Println("Please enter data of Base Rectangle:")
	for trp.BaseRectangle.Length == 0 || trp.BaseRectangle.Width == 0 {
		if trp.BaseRectangle.Length == 0 {
			fmt.Print("Length - ")
			fmt.Scan(&trp.BaseRectangle.Length)
		}
		if trp.BaseRectangle.Width == 0 {
			fmt.Print("Width - ")
			fmt.Scan(&trp.BaseRectangle.Width)
		}
	}
	for trp.Height == 0 {
		fmt.Print("Height - ")
		fmt.Scan(&trp.Height)
	}
	return &trp
}
