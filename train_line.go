package main

type TrainLine struct {
	route []Points
}

func trainLine(relationId uint64) TrainLine {

	//var frankstone_line_id uint64 = 344911

	var points []Points

	frankstone_line := relations[relationId]

	for _, member := range frankstone_line.Member {
		way := ways[member.Ref]
		points = append(points, wayToPoints(way))
	}

	return TrainLine{route: points}
}

//func frankstoneLinePoints() TrainLine {
//	var frankstone_line_id uint64 = 344911
//	return trainLine(frankstone_line_id)
//}

func AllTrainLines() []TrainLine {
	var lines []TrainLine

	for _, relation := range melbourne.Relation {
		if relation.IsTrainRoute() {

			//TODO This line breaks my projections, needs fixing
			if relation.Id != 905345 {
				lines = append(lines, trainLine(relation.Id))
			}

		}
	}

	return lines
}
