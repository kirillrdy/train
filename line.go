package main

import "math"

type Line struct {
	Start Point
	End   Point
}

func (line Line) Dy() float64 {
	return line.End.y - line.Start.y
}

func (line Line) Dx() float64 {
	return line.End.x - line.Start.x
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

const VELOCITY = 1
