package figure

import (
	"math"
	"testing"
)

type figure interface {
	Volume() float64
}

type testDefinition struct {
	name           string
	input          figure
	expectedResult float64
}

func TestVolume(t *testing.T) {
	tests := []testDefinition{
		{
			name:           "Cone",
			input:          &Cone{Radius: 5, Height: 5},
			expectedResult: 130.9,
		},
		{
			name:           "Cube",
			input:          &Cube{Height: 5},
			expectedResult: 125,
		},
		{
			name:           "Cylinder",
			input:          &Cylinder{Height: 5, Radius: 5},
			expectedResult: 392.7,
		},
		{
			name:           "Ellipsoid",
			input:          &Ellipsoid{a: 2, b: 3, c: 4},
			expectedResult: 100.53,
		},
		{
			name:           "Rectangular Prism",
			input:          &RectangularPrism{Length: 3, Width: 2, Height: 5},
			expectedResult: 30,
		},
		{
			name:           "Rectangular Pyramid",
			input:          &RectangularPyramid{&Rectangle{2, 3}, 5},
			expectedResult: 10,
		},
		{
			name:           "Sphere",
			input:          &Sphere{Radius: 5},
			expectedResult: 523.6,
		},
		{
			name:           "TriangularPrism",
			input:          &TriangularPrism{&Triangle{2, 3, 4}, 5},
			expectedResult: 14.52,
		},
		{
			name:           "TriangularPyramid",
			input:          &TriangularPyramid{&Triangle{2, 3, 4}, 5},
			expectedResult: 4.84,
		},
		{
			name:           "TruncatedRectangularPyramid",
			input:          &TruncatedRectangularPyramid{&Rectangle{2, 1}, &Rectangle{4, 3}, 5},
			expectedResult: 31.67,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.expectedResult != math.Round(tt.input.Volume()*100)/100 {
				t.Fatalf("test for %v is failed, expected result - %v, got - %v,\n", tt.name, tt.expectedResult, math.Round(tt.input.Volume()*100)/100)
			}
		})
	}
}
