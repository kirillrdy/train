package model

type Screen struct {
	Width  float64
	Height float64
}

func (screen Screen) ToRectangle() Rectangle {
	return Rectangle{Point{0, 0}, Point{screen.Width, screen.Height}}
}
