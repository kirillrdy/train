package model

import (
	"fmt"
	"math/rand"

	"github.com/kirillrdy/nadeshiko"
	"github.com/kirillrdy/nadeshiko/html"
	"github.com/sparkymat/webdsl/css"
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

func (train Train) Selector() string {
	return css.Id(train.Id).Selector()
}

func (train Train) AppendToPage(document nadeshiko.Document) {
	document.JQuery(css.Body).Append(html.P().Id(train.Id).Text(")"))
	document.JQuery(train).SetCss("position", "absolute")
}

func (train Train) RemoveFromPage(document nadeshiko.Document) {
	document.JQuery(train).Remove()
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

func (train Train) Draw(document nadeshiko.Document) {
	document.JQuery(train).SetCss("left", fmt.Sprintf("%fpx", train.Position.X+2))
	document.JQuery(train).SetCss("top", fmt.Sprintf("%fpx", train.Position.Y-17))
}
