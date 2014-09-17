package main

import (
	"fmt"
	"strings"

	"github.com/kirillrdy/nadeshiko"
	"github.com/kirillrdy/train/model"
)

var projection model.Projection
var city model.City
var svg_element string

func setUp() {
	city := model.LoadCity("melbourne.json")

	scren := model.Rectangle{model.Point{0, 0}, model.Point{1600, 1200}}
	fake_world := model.Rectangle{model.Point{144.5265, -37.6474}, model.Point{145.6032, -38.1427}}
	projection := model.Projection{Original: fake_world, Destination: scren}

	svg_element = `<svg width="1600" height="1200" >
				%s
			</svg> `

	var svg_paths []string

	for _, line := range city.TrainLines {
		for _, section := range line.Sections {
			svg_paths = append(svg_paths, section.Points.Translate(projection).ToSvgPath())
		}
	}

	svg_element = fmt.Sprintf(svg_element, strings.Join(svg_paths, "\n"))

}

func AddMap(document nadeshiko.Document) {
	document.JQuery("body").Append(svg_element)
}

func handler(document nadeshiko.Document) {

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
