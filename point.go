package main

import (
	"fmt"
	"math"
	"strings"
)

type Point struct {
	x, y float64
}

type Points []Point

func (points Points) pointsSvgToPath() string {
	var points_as_string []string
	for _, point := range points {
		points_as_string = append(points_as_string, fmt.Sprintf("%f %f", point.x, point.y))
	}
	return fmt.Sprintf("M %s", strings.Join(points_as_string, " L "))
}

func (point Point) Transform(direction float64, velocity float64) Point {
	new_location := Point{}
	new_location.x = (velocity * math.Cos(direction)) + point.x
	new_location.y = (velocity * math.Sin(direction)) + point.y
	return new_location
}

func (point Point) Floored() Point {
	floored := Point{}
	floored.x = math.Floor(point.x)
	floored.y = math.Floor(point.y)
	return floored
}

func (point Point) DistanceTo(second_point Point) float64 {
	xdiff := point.x - second_point.x
	ydiff := point.y - second_point.y
	return math.Sqrt((xdiff * xdiff) + (ydiff * ydiff))
}
