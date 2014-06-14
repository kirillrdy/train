package main

import (
	"fmt"
	"strings"
)

type Point struct {
	x, y float64
}

func pointsToPath(points []Point) string {
	var points_as_string []string
	for _, point := range points {
		points_as_string = append(points_as_string, fmt.Sprintf("%f %f", point.x, point.y))
	}
	return fmt.Sprintf("M %s", strings.Join(points_as_string, " L "))
}
