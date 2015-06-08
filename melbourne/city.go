package melbourne

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/kirillrdy/nadeshiko/html"
	"github.com/kirillrdy/train/model"
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

func RailMap(projection model.Projection) html.Node {

	var svg_paths []string

	for _, line := range city.TrainLines {
		for _, section := range line.Sections {
			svg_paths = append(svg_paths, section.Points.Translate(projection).ToSvgPath())
		}
	}

	svg_element := html.Svg().WidthFloat(projection.Destination.Max.X).HeightFloat(projection.Destination.Max.Y)
	svg_element = svg_element.TextUnsafe(fmt.Sprintf(strings.Join(svg_paths, "\n")))
	return svg_element
}
