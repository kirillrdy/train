package main

import (
	"fmt"
	"math/rand"

	"github.com/kirillrdy/nadeshiko"
)

type Train struct {
	Id       string
	Position Point
}

func NewTrain() Train {
	random := rand.Int63()
	return Train{Id: fmt.Sprintf("%d", random)}
}

func (train Train) CssSelector() string {
	return "#" + train.Id
}

func (train Train) AppendToPage(conneciton *nadeshiko.Connection) {
	conneciton.JQuery("body").Append(fmt.Sprintf("<p id='%s'>Hello</p>", train.Id))
	conneciton.JQuery(train.CssSelector()).SetCss("position", "absolute")
}

func (train *Train) Step(point Point) {
	direction_vector := Line{train.Position, point}
	fmt.Printf("Facing direction %f\n", direction_vector.AngleDegrees())
	train.Position = train.Position.Transform(direction_vector.AngleRadians(), VELOCITY)
}

func (train Train) At(point Point) bool {
	val := train.Position.DistanceTo(point) < 0.5
	fmt.Printf("%b %v %v \n", val, train.Position.Floored(), point.Floored())
	return val
}

func (train Train) Draw(conneciton *nadeshiko.Connection) {
	selector := "#" + train.Id
	conneciton.JQuery(selector).SetCss("left", fmt.Sprintf("%fpx", train.Position.x))
	conneciton.JQuery(selector).SetCss("top", fmt.Sprintf("%fpx", train.Position.y))
}
