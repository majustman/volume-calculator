package figure

type Rectangle struct {
	Length float64
	Width  float64
}

func (r *Rectangle) Area() float64 {
	return r.Length * r.Width
}
