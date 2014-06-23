package model

import "math"

type Point struct {
	X, Y float64
}

func (point Point) Transform(direction float64, velocity float64) Point {
	new_location := Point{}
	new_location.X = (velocity * math.Cos(direction)) + point.X
	new_location.Y = (velocity * math.Sin(direction)) + point.Y
	return new_location
}

func (point Point) Floored() Point {
	floored := Point{}
	floored.X = math.Floor(point.X)
	floored.Y = math.Floor(point.Y)
	return floored
}

func (point Point) DistanceTo(second_point Point) float64 {
	xdiff := point.X - second_point.X
	ydiff := point.Y - second_point.Y
	return math.Sqrt((xdiff * xdiff) + (ydiff * ydiff))
}
