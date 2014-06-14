package main

import (
	"fmt"
	"time"

	"github.com/kirillrdy/nadeshiko"
)

type MainActivity struct {
}

func (activity MainActivity) Start(conneciton *nadeshiko.Connection) {

	points := []Point{
		Point{100, 10},
		Point{300, 300},
		Point{10, 400},
	}

	data := `<svg height="480" width="640">
			<path d="%s" fill="none" stroke="blue"/>
			</svg> `

	conneciton.JQuery("body").Append(fmt.Sprintf(data, pointsToPath(points)))

	train := Train{}
	train.AppendToPage(conneciton)

	current_index := 0

	go func() {
		for {

			train.Step(points[current_index])
			train.Draw(conneciton)
			time.Sleep(100 * time.Millisecond)

			if train.At(points[current_index]) {
				current_index += 1
			}
		}

	}()
}
