package main

import (
	"github.com/kirillrdy/osm"
	"github.com/kirillrdy/train/model"
)

func wayToSection(osm osm.Osm, way osm.Way) (section model.Section) {

	var points model.Points

	for _, nd := range way.Nd {
		node := osm.NodeById(nd.Ref)
		if node != nil {
			point := model.Point{node.Lon, node.Lat}
			points = append(points, point)
		}

	}
	//TODO extract stations info from way.Nds
	section.Points = points
	return section
}

func osmRelationToTrainLine(osm osm.Osm, relation osm.Relation) (trainLine model.TrainLine) {

	var sections model.Sections

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

func AllTrainLines(osm osm.Osm) (lines model.TrainLines) {

	for _, relation := range TrainLinesRelations(osm) {
		lines = append(lines, osmRelationToTrainLine(osm, relation))
	}

	return lines
}

func main() {
	melbourneOsm := osm.LoadPackagedMelbourne()
	melbourneOsm.BuildIndex()
	melbourne := model.City{}

	trainLines := AllTrainLines(*melbourneOsm)
	melbourne.TrainLines = trainLines
	melbourne.Save("melbourne.json")
}
