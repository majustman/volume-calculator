package service

import (
	"fmt"
	"github.com/majustman/volume-calculator/figure"
	"math"
)

const (
	cone = iota + 1
	cube
	cylinder
	ellipsoid
	rectangularPrism
	rectangularPyramid
	sphere
	triangularPrism
	triangularPyramid
	truncatedRectangularPyramid
)

func Run() {
	i := 0
	for i == 0 {
		fmt.Println("Please choose on of the following figures by entering the number:\n" +
			"1) Cone\n" +
			"2) Cube\n" +
			"3) Cylinder\n" +
			"4) Ellipsoid\n" +
			"5) RectangularPrism\n" +
			"6) RectangularPyramid\n" +
			"7) Sphere\n" +
			"8) TriangularPrism\n" +
			"9) TriangularPyramid\n" +
			"10) TruncatedRectangularPyramid")
		fmt.Scan(&i)
		if i < 1 || i > 10 {
			i = 0
		}
	}
	var v float64
	switch i {
	case cone:
		v = math.Round(figure.NewCone().Volume()*100) / 100
	case cube:
		v = math.Round(figure.NewCube().Volume()*100) / 100
	case cylinder:
		v = math.Round(figure.NewCylinder().Volume()*100) / 100
	case ellipsoid:
		v = math.Round(figure.NewEllipsoid().Volume()*100) / 100
	case rectangularPrism:
		v = math.Round(figure.NewRectangularPrism().Volume()*100) / 100
	case rectangularPyramid:
		v = math.Round(figure.NewRectangularPyramid().Volume()*100) / 100
	case sphere:
		v = math.Round(figure.NewSphere().Volume()*100) / 100
	case triangularPrism:
		v = math.Round(figure.NewTriangularPrism().Volume()*100) / 100
	case triangularPyramid:
		v = math.Round(figure.NewTriangularPyramid().Volume()*100) / 100
	case truncatedRectangularPyramid:
		v = math.Round(figure.NewTruncatedRectangularPyramid().Volume()*100) / 100
	}
	fmt.Printf("The volume equal to %v\n", v)
}
