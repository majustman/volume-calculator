package figure

import "fmt"

type TriangularPyramid struct {
	*Triangle
	Height float64
}

func (p *TriangularPyramid) Volume() float64 {
	return (p.Triangle.Area() * p.Height) / 3
}

func NewTriangularPyramid() *TriangularPyramid {
	var tp TriangularPyramid
	tp.Triangle = &Triangle{}
	fmt.Println("Please enter data:")
	for tp.SideA == 0 || tp.SideB == 0 || tp.SideC == 0 || tp.Height == 0 {
		if tp.SideA == 0 {
			fmt.Print("Side A of base triangle - ")
			fmt.Scan(&tp.SideA)
		}
		if tp.SideB == 0 {
			fmt.Print("Side B of base triangle - ")
			fmt.Scan(&tp.SideB)
		}
		if tp.SideC == 0 {
			fmt.Print("Side C of base triangle - ")
			fmt.Scan(&tp.SideC)
		}
		if tp.Height == 0 {
			fmt.Print("Height - ")
			fmt.Scan(&tp.Height)
		}
	}
	return &tp
}
