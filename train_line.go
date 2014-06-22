package main

import "github.com/kirillrdy/osm"

type TrainLine struct {
	route []Points
}

func trainLine(relationId uint64) TrainLine {

	var points []Points

	frankstone_line := relationsCache[relationId]

	for _, member := range frankstone_line.Member {
		way := waysCache[member.Ref]
		points = append(points, wayToPoints(way))
	}

	return TrainLine{route: points}
}

//func frankstoneLinePoints() TrainLine {
//	var frankstone_line_id uint64 = 344911
//	return trainLine(frankstone_line_id)
//}

func TrainLinesRelations() []osm.Relation {
	var results []osm.Relation
	for _, relation := range melbourne.Relation {
		if relation.IsTrainRoute() {

			//TODO This line breaks my projections, needs fixing
			if relation.Id != 905345 {
				results = append(results, relation)
			}
		}
	}
	return results

}

func RelationsToWays(relations []osm.Relation) []osm.Way {
	var ways []osm.Way
	for _, relation := range relations {
		for _, member := range relation.Member {
			ways = append(ways, waysCache[member.Ref])
		}

	}
	return ways
}

func WaysToNodes(ways []osm.Way) []osm.Node {
	var nodes []osm.Node
	for _, way := range ways {
		for _, ng := range way.Nd {
			nodes = append(nodes, nodesCache[ng.Ref])
		}
	}
	return nodes
}

func AllTrainLines() []TrainLine {
	var lines []TrainLine

	for _, relation := range TrainLinesRelations() {
		lines = append(lines, trainLine(relation.Id))
	}

	return lines
}
