package main

import (
	"github.com/kirillrdy/nadeshiko"
	"github.com/kirillrdy/train/lib"
)

func main() {
	nadeshiko.Nadeshiko("/", lib.Handler)
	nadeshiko.Start()
}
