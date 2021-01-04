package interfaces

type Shape interface {
	Area() float64
	Perimeter() float64
}

type React struct {
	Width float64
	Hight float64
}

func (react React) Area() float64 {
	return react.Hight * react.Width
}

func (react React) Perimeter() float64 {
	return 2*react.Hight * react.Width
}
