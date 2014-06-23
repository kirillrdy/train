package model

import "github.com/kirillrdy/osm"

type TrainLine struct {
	Name     string
	Sections Sections
}

type TrainLines []TrainLine

func osmRelationToTrainLine(osm osm.Osm, relation osm.Relation) (trainLine TrainLine) {

	var sections Sections

	for _, member := range relation.WayMembers() {
		way := osm.WayById(member.Ref)
		if way != nil {
			sections = append(sections, wayToSection(osm, *way))
		}
	}

	trainLine.Sections = sections
	return trainLine
}

//func frankstoneLinePoints() TrainLine {
//	var frankstone_line_id uint64 = 344911
//	return trainLine(frankstone_line_id)
//}

func TrainLinesRelations(osm osm.Osm) (results []osm.Relation) {
	for _, relation := range osm.Relation {
		if relation.IsTrainRoute() {

			//TODO This line breaks my projections, needs fixing
			if relation.Id != 905345 {
				results = append(results, relation)
			}
		}
	}
	return results

}

func AllTrainLines(osm osm.Osm) (lines TrainLines) {

	for _, relation := range TrainLinesRelations(osm) {
		lines = append(lines, osmRelationToTrainLine(osm, relation))
	}

	return lines
}
