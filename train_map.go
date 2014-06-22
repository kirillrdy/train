package main

import "github.com/kirillrdy/osm"

func wayToPoints(way osm.Way) Points {

	var points Points

	for _, nd := range way.Nd {
		node := nodes[nd.Ref]
		point := Point{node.Lon, node.Lat}
		points = append(points, point)

	}
	return points
}
