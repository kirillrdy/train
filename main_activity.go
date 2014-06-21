package main

import (
	"fmt"
	"time"

	"github.com/kirillrdy/nadeshiko"
)

type MainActivity struct {
	trains []*Train
}

func (activity MainActivity) Start(conneciton *nadeshiko.Connection) {

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

	data := `<svg height="480" width="640">
			<path d="%s" fill="none" stroke="blue"/>
			</svg> `

	conneciton.JQuery("body").Append(fmt.Sprintf(data, Points(path).pointsSvgToPath()))

	button := `<input type="button" id="add" value="Add Train">`
	conneciton.JQuery("body").Append(button)

	conneciton.JQuery("#add").Click(func() {
		//activity.addTrainToPath(path, conneciton)
		activity.addAtrain(path, conneciton)
	})

	go func() {
		for {
			time.Sleep(10 * time.Millisecond)
			for _, train := range activity.trains {
				train.OneFrame()
			}

			conneciton.StartBuffer()

			for _, train := range activity.trains {
				train.Draw(conneciton)
			}

			conneciton.FlushBuffer()

			var non_dead_head []*Train

			for _, train := range activity.trains {
				if train.ShouldBeRemoved() {
					//TODO also remove from trains
					train.RemoveFromPage(conneciton)
				} else {
					non_dead_head = append(non_dead_head, train)
				}
			}
			activity.trains = non_dead_head
		}
	}()

	//go func() {
	//	for {
	//		addTrainToPath(path, conneciton)
	//		time.Sleep(3000 * time.Millisecond)
	//	}
	//}()

}

func (activity *MainActivity) addAtrain(path []Point, conneciton *nadeshiko.Connection) {
	for i := 0; i < 5; i++ {
		activity.addTrainToPath(path, conneciton)
		time.Sleep(120 * time.Millisecond)
	}
}

func (activity *MainActivity) addTrainToPath(path []Point, conneciton *nadeshiko.Connection) {

	train := NewTrain(path)
	train.AppendToPage(conneciton)
	activity.trains = append(activity.trains, &train)

}
