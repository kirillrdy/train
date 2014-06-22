package main

import "github.com/kirillrdy/nadeshiko"
import "github.com/kirillrdy/osm"

var melbourne = osm.LoadFromJson()

var nodes = map[uint64]osm.Node{}
var ways = map[uint64]osm.Way{}
var relations = map[uint64]osm.Relation{}

func main() {

	for _, node := range melbourne.Node {
		nodes[node.Id] = node
	}

	for _, way := range melbourne.Way {
		ways[way.Id] = way
	}

	for _, relation := range melbourne.Relation {
		relations[relation.Id] = relation
	}

	routes := nadeshiko.Routes{}
	routes.Activity("/", NewMainActivity())
	nadeshiko.Start(routes)
}
