package main

import (
	"fmt"
	"time"

	"github.com/kirillrdy/nadeshiko"
)

type MainActivity struct {
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
	}

	data := `<svg height="480" width="640">
			<path d="%s" fill="none" stroke="blue"/>
			</svg> `

	conneciton.JQuery("body").Append(fmt.Sprintf(data, Points(path).pointsSvgToPath()))

	button := `<input type="button" id="add" value="Add Train">`
	conneciton.JQuery("body").Append(button)

	conneciton.JQuery("#add").Click(func() {
		//addTrainToPath(path, conneciton)
		addAtrain(path, conneciton)
	})

	//go func() {
	//	for {
	//		addTrainToPath(path, conneciton)
	//		time.Sleep(3000 * time.Millisecond)
	//	}
	//}()

}

func addAtrain(path []Point, conneciton *nadeshiko.Connection) {
	for i := 0; i < 5; i++ {
		addTrainToPath(path, conneciton)
		time.Sleep(120 * time.Millisecond)
	}
}

func addTrainToPath(path []Point, conneciton *nadeshiko.Connection) {

	train := NewTrain()
	train.AppendToPage(conneciton)
	current_index := 0

	go func() {
		for {

			current_target_point := path[current_index]

			train.Step(current_target_point)
			train.Draw(conneciton)
			time.Sleep(10 * time.Millisecond)

			if train.At(current_target_point) {
				current_index += 1
				if current_index == len(path) {
					train.RemoveFromPage(conneciton)
					return
				}
			}
		}

	}()
}
