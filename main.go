package main

import (
	"github.com/kirillrdy/nadeshiko"
	"github.com/kirillrdy/train/model"
)

func main() {

	nadeshiko.StartActivity(NewMainActivity(model.LoadCity("melbourne.json")))
	nadeshiko.Start()
}
