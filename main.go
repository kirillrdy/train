package main

import (
	"github.com/kirillrdy/nadeshiko"
	"github.com/kirillrdy/train/model"
)

func main() {

	routes := nadeshiko.Routes{}
	routes.Activity("/", NewMainActivity(model.LoadCity("melbourne.json")))
	nadeshiko.Start(routes)
}
