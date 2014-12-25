package main

import (
	"time"

	"github.com/kirillrdy/nadeshiko"
	"github.com/kirillrdy/nadeshiko/html"
	"github.com/sparkymat/webdsl/css"
)

var trains []*Train

func MainActivity(document *nadeshiko.Document) {

	path := []Point{
		Point{0, 0},
		Point{30, 0},
		Point{30, 40},
		Point{200, 70},
		Point{300, 50},
		Point{330, 90},
		Point{160, 200},
		Point{200, 250},
		Point{300, 300},
		Point{10, 400},
		Point{100, 250},
		Point{400, 150},
	}

	svg := html.Svg().Height(480).Width(640).Children(
		html.Path().Attribute("fill", "none").Attribute("stroke", "blue").Attribute("d", Points(path).pointsSvgToPath()),
	)

	document.JQuery(css.Body).Append(svg)

	buttonId := css.Id("add")

	button := html.Input().Type("button").Id(buttonId).Value("Add Train")
	document.JQuery(css.Body).Append(button)

	document.JQuery(buttonId).Click(func() {
		//activity.addTrainToPath(path, conneciton)
		addAtrain(path, document)
	})

	go func() {
		for {
			time.Sleep(10 * time.Millisecond)
			for _, train := range trains {
				train.OneFrame()
			}

			document.StartBuffer()

			for _, train := range trains {
				train.Draw(document)
			}

			document.FlushBuffer()

			var non_dead_head []*Train

			for _, train := range trains {
				if train.ShouldBeRemoved() {
					//TODO also remove from trains
					train.RemoveFromPage(document)
				} else {
					non_dead_head = append(non_dead_head, train)
				}
			}
			trains = non_dead_head
		}
	}()

	//go func() {
	//	for {
	//		addTrainToPath(path, conneciton)
	//		time.Sleep(3000 * time.Millisecond)
	//	}
	//}()

}

func addAtrain(path []Point, document *nadeshiko.Document) {
	for i := 0; i < 100; i++ {
		addTrainToPath(path, document)
		time.Sleep(120 * time.Millisecond)
	}
}

func addTrainToPath(path []Point, document *nadeshiko.Document) {

	train := NewTrain(path)
	train.AppendToPage(document)
	trains = append(trains, &train)

}
