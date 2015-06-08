package model

import "math"

type Line struct {
	Start Point
	End   Point
}

func (line Line) Dy() float64 {
	return line.End.Y - line.Start.Y
}

func (line Line) Dx() float64 {
	return line.End.X - line.Start.X
}

//Only useful for printing
func (line Line) AngleDegrees() float64 {
	return line.AngleRadians() * 180 / math.Pi
}

func (line Line) Angle() float64 {
	return line.AngleRadians()
}

func (line Line) AngleRadians() float64 {
	theta := math.Atan2(line.Dy(), line.Dx())
	return theta
}
