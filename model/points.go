package model

import (
	"fmt"
	"strings"
)

type Points []Point

func (points Points) Translate(projection Projection) Points {
	var points_translated Points
	for _, point := range points {
		points_translated = append(points_translated, projection.Translate(point))
	}
	return points_translated
}

func (points Points) ToSvgPath() string {
	if len(points) == 0 {
		return ""
	}

	var points_as_string []string
	for _, point := range points {
		points_as_string = append(points_as_string, fmt.Sprintf("%f %f", point.X, point.Y))
	}
	path := fmt.Sprintf("M %s", strings.Join(points_as_string, " L "))
	return fmt.Sprintf(`<path d="%s" fill="none" stroke="blue"/>`, path)
}
