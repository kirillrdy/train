package main

import (
	"fmt"
	"math/rand"

	"github.com/kirillrdy/nadeshiko"
)

type Train struct {
	Id                string
	Position          Point
	currentPointIndex int
	path              []Point
}

func NewTrain(path []Point) Train {
	random := rand.Int63()
	return Train{Id: fmt.Sprintf("%d", random),
		path:              path,
		currentPointIndex: 1,
		Position:          path[0]}
}

func (train Train) CssSelector() string {
	return "#" + train.Id
}

func (train Train) AppendToPage(conneciton *nadeshiko.Connection) {
	conneciton.JQuery("body").Append(fmt.Sprintf("<p id='%s'>O</p>", train.Id))
	conneciton.JQuery(train.CssSelector()).SetCss("position", "absolute")
}

func (train Train) RemoveFromPage(conneciton *nadeshiko.Connection) {
	conneciton.JQuery(train.CssSelector()).Remove()
}

func (train *Train) Step(point Point) {
	direction_vector := Line{train.Position, point}
	//fmt.Printf("Facing direction %f\n", direction_vector.AngleDegrees())
	newPosition := train.Position.Transform(direction_vector.AngleRadians(), VELOCITY)
	train.Position = newPosition
}

func (train Train) CurrentTragetPoint() Point {
	if train.ShouldBeRemoved() {
		return train.path[len(train.path)-1]
	} else {
		return train.path[train.currentPointIndex]
	}
}

func (train Train) ShouldBeRemoved() bool {
	return train.currentPointIndex >= len(train.path)
}

func (train *Train) OneFrame() {

	target := train.CurrentTragetPoint()

	train.Step(target)

	if train.At(target) {
		train.currentPointIndex += 1
	}
}

func (train Train) At(point Point) bool {
	val := train.Position.DistanceTo(point) < 0.5
	//fmt.Printf("%b %v %v \n", val, train.Position.Floored(), point.Floored())
	return val
}

func (train Train) Draw(conneciton *nadeshiko.Connection) {
	selector := "#" + train.Id
	conneciton.JQuery(selector).SetCss("left", fmt.Sprintf("%fpx", train.Position.x+2))
	conneciton.JQuery(selector).SetCss("top", fmt.Sprintf("%fpx", train.Position.y-17))
}
