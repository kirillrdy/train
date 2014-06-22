package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/kirillrdy/nadeshiko"
)

type MainActivity struct {
	trains     []*Train
	points     []Point
	projection Projection
}

func NewMainActivity() MainActivity {

	scren := Rectangle{Point{0, 0}, Point{1600, 1200}}
	fake_world := Rectangle{Point{144.5265, -37.6474}, Point{145.6032, -38.1427}}
	projection := Projection{Original: fake_world, Destination: scren}

	return MainActivity{projection: projection}
}

func (activity MainActivity) AddMap(conneciton *nadeshiko.Connection) {

	svg_element := `<svg width="1600" height="1200" >
				%s
			</svg> `

	var svg_paths []string

	for _, line := range AllTrainLines() {
		for _, points := range line.route {
			svg_paths = append(svg_paths, points.Translate(activity.projection).ToSvgPath())
		}
	}

	allPaths := fmt.Sprintf(svg_element, strings.Join(svg_paths, "\n"))

	conneciton.JQuery("body").Append(allPaths)
}

func (activity MainActivity) Start(conneciton *nadeshiko.Connection) {

	activity.AddMap(conneciton)

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

func (activity *MainActivity) addAtrain(path []Point, conneciton *nadeshiko.Connection) {
	for i := 0; i < 100; i++ {
		activity.addTrainToPath(path, conneciton)
		time.Sleep(120 * time.Millisecond)
	}
}

func (activity *MainActivity) addTrainToPath(path []Point, conneciton *nadeshiko.Connection) {

	train := NewTrain(path)
	train.AppendToPage(conneciton)
	activity.trains = append(activity.trains, &train)

}
