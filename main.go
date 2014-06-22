package main

import "github.com/kirillrdy/nadeshiko"
import "github.com/kirillrdy/osm"

var melbourne = osm.LoadPackagedMelbourne()

var nodesCache = map[uint64]osm.Node{}
var waysCache = map[uint64]osm.Way{}
var relationsCache = map[uint64]osm.Relation{}

func main() {

	for _, node := range melbourne.Node {
		nodesCache[node.Id] = node
	}

	for _, way := range melbourne.Way {
		waysCache[way.Id] = way
	}

	for _, relation := range melbourne.Relation {
		relationsCache[relation.Id] = relation
	}

	smallerOsm := osm.Osm{}
	smallerOsm.Relation = TrainLinesRelations()
	smallerOsm.Way = RelationsToWays(smallerOsm.Relation)
	smallerOsm.Node = WaysToNodes(smallerOsm.Way)

	routes := nadeshiko.Routes{}
	routes.Activity("/", NewMainActivity())
	nadeshiko.Start(routes)
}
