package melbourne

import (
	"log"
	"time"

	"github.com/kirillrdy/nadeshiko/html"

	"github.com/kirillrdy/train/model"
	"github.com/sparkymat/webdsl/css"
)

var city model.City

//var Boundaries model.Rectangle = model.Rectangle{model.Point{144.5265, -37.6474}, model.Point{145.6032, -38.1427}}
var Boundaries model.Rectangle = model.Rectangle{
	model.Point{144.6302, -37.6474},
	model.Point{145.3240, -38.0052},
}

func init() {
	now := time.Now()
	city = model.LoadCity("melbourne.json")
	log.Printf("Loading melbourne.json took %s", time.Since(now))
}

func RailMap(svgId css.Id, projection model.Projection) []html.Node {

	var svg_paths []html.Node

	for _, line := range city.TrainLines {
		for _, section := range line.Sections {
			svg_paths = append(svg_paths, section.Points.Translate(projection).ToSvgPath())
		}
	}

	return svg_paths
}
