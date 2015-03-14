package main

import (
	"github.com/kirillrdy/nadeshiko"
	"github.com/kirillrdy/train/lib"
)

func main() {
	lib.SetUp()
	nadeshiko.Nadeshiko("/", lib.Handler)
	nadeshiko.Start()
}
