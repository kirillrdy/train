package main

import "github.com/kirillrdy/osm"

type TrainLine struct {
	route []Points
}

func trainLine(relationId uint64) TrainLine {

	var points []Points

	frankstone_line := melbourne.RelationById(relationId)

	for _, member := range frankstone_line.Member {
		way := melbourne.WayById(member.Ref)
		if way != nil {
			points = append(points, wayToPoints(*way))
		}
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

//TODO relation members also sometimes contain nodes
func RelationsToWays(relations []osm.Relation) []osm.Way {
	var ways []osm.Way
	for _, relation := range relations {
		for _, member := range relation.WayMembers() {
			way := melbourne.WayById(member.Ref)
			if way != nil {
				ways = append(ways, *way)
			}
		}
	}
	return ways
}

func WaysToNodes(ways []osm.Way) []osm.Node {
	var nodes []osm.Node
	for _, way := range ways {
		for _, ng := range way.Nd {
			node := melbourne.NodeById(ng.Ref)
			if node != nil {
				nodes = append(nodes, *node)
			}

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
