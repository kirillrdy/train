package main

import (
	"github.com/kirillrdy/nadeshiko"
	"github.com/kirillrdy/train/lib"
	"log"
	"time"
)

//TODO move to util
func Bench(name string, function func()) {
	start := time.Now()
	function()
	log.Printf("%#s taken %s\n", name, time.Since(start))
}

func main() {
	Bench("SetUp", lib.SetUp)
	nadeshiko.Nadeshiko("/", lib.Handler)
	nadeshiko.Start()
}
