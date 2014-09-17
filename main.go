package main

import "github.com/kirillrdy/nadeshiko"

func main() {
	setUp()
	nadeshiko.Nadeshiko("/", handler)
	nadeshiko.Start()
}
