package main

import "github.com/kirillrdy/nadeshiko"

func main() {
	routes := nadeshiko.Routes{}
	routes.Activity("/", MainActivity{})
	nadeshiko.Start(routes)
}
