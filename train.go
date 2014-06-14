package main

import (
	"fmt"

	"github.com/kirillrdy/nadeshiko"
)

type Train struct {
	x, y float64
}

func (train Train) AppendToPage(conneciton *nadeshiko.Connection) {
	conneciton.JQuery("body").Append("<p id='train'>Hello</p>")
	conneciton.JQuery("#train").SetCss("position", "absolute")
}

func (train *Train) Step(point Point) {
	if train.x < point.x {
		train.x += 1
		x_diff := (point.x - train.x)
		if x_diff != 0 {
			delta := (point.y - train.y) / x_diff
			train.y += delta
		}
	}
}

func (train Train) At(point Point) bool {
	val := train.x == point.x && train.y == point.y
	fmt.Printf("%b %f %f %f %f  \n", val, train.x, point.x, train.y, point.y)
	return val
}

func (train *Train) Move(point Point) {
	train.x = point.x
	train.y = point.y
}

func (train Train) Draw(conneciton *nadeshiko.Connection) {
	conneciton.JQuery("#train").SetCss("left", fmt.Sprintf("%fpx", train.x))
	conneciton.JQuery("#train").SetCss("top", fmt.Sprintf("%fpx", train.y))
}
