package model

import "github.com/kirillrdy/osm"

type Section struct {
	Points Points
	//Stations Stations
}

type Sections []Section

func wayToSection(osm osm.Osm, way osm.Way) (section Section) {

	var points Points

	for _, nd := range way.Nd {
		node := osm.NodeById(nd.Ref)
		if node != nil {
			point := Point{node.Lon, node.Lat}
			points = append(points, point)
		}

	}
	//TODO extract stations info from way.Nds
	section.Points = points
	return section
}
