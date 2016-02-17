package lib

import (
	"fmt"
	"strings"

	"github.com/kirillrdy/nadeshiko"
	"github.com/kirillrdy/nadeshiko/html"
	"github.com/kirillrdy/train/model"
	"github.com/sparkymat/webdsl/css"
)

var projection model.Projection
var city model.City
var svgElement html.Node

//SetUp loads all required data and sets whatever needed
//TODO possibly move to init ?
func SetUp() {
	city := model.LoadCity("melbourne.json")

	screen := model.Rectangle{Min: model.Point{X: 0, Y: 0}, Max: model.Point{X: 1600, Y: 1200}}
	melbourneBoundary := model.Rectangle{Min: model.Point{X: 144.5265, Y: -37.6474}, Max: model.Point{X: 145.6032, Y: -38.1427}}
	projection := model.Projection{Original: melbourneBoundary, Destination: screen}

	var svgPaths []string

	for _, line := range city.TrainLines {
		for _, section := range line.Sections {
			svgPaths = append(svgPaths, section.Points.Translate(projection).ToSvgPath())
		}
	}

	svgElement = html.Svg().Height(1200).Width(1600).TextUnsafe(fmt.Sprintf(strings.Join(svgPaths, "\n")))
}

func AddMap(document *nadeshiko.Document) {
	document.JQuery(css.Body).Append(svgElement)
}

func Handler(document *nadeshiko.Document) {

	AddMap(document)

	//button := `<input type="button" id="add" value="Add Train">`
	//conneciton.JQuery("body").Append(button)

	//conneciton.JQuery("#add").Click(func() {
	//	limit += 1
	//	conneciton.JQuery("svg").Remove()
	//	activity.AddMap(conneciton, limit)
	//})

	//go func() {
	//	for {
	//		time.Sleep(10 * time.Millisecond)
	//		for _, train := range activity.trains {
	//			train.OneFrame()
	//		}

	//		conneciton.StartBuffer()

	//		for _, train := range activity.trains {
	//			train.Draw(conneciton)
	//		}

	//		conneciton.FlushBuffer()

	//		var non_dead_head []*Train

	//		for _, train := range activity.trains {
	//			if train.ShouldBeRemoved() {
	//				train.RemoveFromPage(conneciton)
	//			} else {
	//				non_dead_head = append(non_dead_head, train)
	//			}
	//		}
	//		activity.trains = non_dead_head
	//	}
	//}()

}

//func (activity *MainActivity) addAtrain(path []Point, conneciton *nadeshiko.Connection) {
//	for i := 0; i < 100; i++ {
//		activity.addTrainToPath(path, conneciton)
//		time.Sleep(120 * time.Millisecond)
//	}
//}
//
//func (activity *MainActivity) addTrainToPath(path []Point, conneciton *nadeshiko.Connection) {
//
//	train := NewTrain(path)
//	train.AppendToPage(conneciton)
//	activity.trains = append(activity.trains, &train)
//
//}
