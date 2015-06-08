package main

import (
	"fmt"
	"log"
	"time"

	"github.com/kirillrdy/osm"
	"github.com/kirillrdy/train/model"
)

func wayToSection(osm *osm.Osm, index *osm.OsmIndex, way osm.Way) (section model.Section) {

	var points model.Points

	for _, nd := range way.Nd {
		node := index.NodeById(nd.Ref)
		if node != nil {
			point := model.Point{node.Lon, node.Lat}
			points = append(points, point)
		}

	}
	//TODO extract stations info from way.Nds
	section.Points = points
	return section
}

func osmRelationToTrainLine(osm *osm.Osm, index *osm.OsmIndex, relation osm.Relation) (trainLine model.TrainLine) {

	var sections model.Sections

	for _, member := range relation.WayMembers() {
		way := index.WayById(member.Ref)
		if way != nil {
			sections = append(sections, wayToSection(osm, index, *way))
		}
	}

	trainLine.Sections = sections
	return trainLine
}

//func frankstoneLinePoints() TrainLine {
//	var frankstone_line_id uint64 = 344911
//	return trainLine(frankstone_line_id)
//}

func trainLinesRelations(osm *osm.Osm) (results []osm.Relation) {
	for i := range osm.Relation {
		relation := osm.Relation[i]
		if relation.IsTrainRoute() {

			//TODO This line breaks my projections, needs fixing
			if relation.Id != 905345 {
				results = append(results, *relation)
			}
		}
	}
	return results

}

func AllTrainLines(osm *osm.Osm) (lines model.TrainLines) {
	index := osm.BuildIndex()

	for _, relation := range trainLinesRelations(osm) {
		lines = append(lines, osmRelationToTrainLine(osm, &index, relation))
	}

	return lines
}

func main() {
	log.Print("Loading OSM data for melbourne")

	go func() {
		ticker := time.Tick(time.Second)
		for range ticker {
			fmt.Print(".")
		}
	}()

	melbourneOsm := osm.LoadPackagedMelbourne()
	log.Print("Building Index")

	melbourne := model.City{}

	log.Print("Extracting train lines")
	trainLines := AllTrainLines(melbourneOsm)
	melbourne.TrainLines = trainLines
	log.Print("Saving melbourne.json")
	melbourne.Save("melbourne.json")
}
