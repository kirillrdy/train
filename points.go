package main

import (
	"fmt"
	"strings"

	"github.com/kirillrdy/osm"
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
		points_as_string = append(points_as_string, fmt.Sprintf("%f %f", point.x, point.y))
	}
	path := fmt.Sprintf("M %s", strings.Join(points_as_string, " L "))
	return fmt.Sprintf(`<path d="%s" fill="none" stroke="blue"/>`, path)
}

func wayToPoints(way osm.Way) Points {

	var points Points

	for _, nd := range way.Nd {
		node := nodesCache[nd.Ref]
		point := Point{node.Lon, node.Lat}
		points = append(points, point)

	}
	return points
}
