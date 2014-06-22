package main

import "github.com/kirillrdy/nadeshiko"
import "github.com/kirillrdy/osm"

var melbourne = osm.LoadPackagedMelbourne()

func main() {

	//smallerOsm := osm.Osm{}
	//smallerOsm.Relation = TrainLinesRelations()
	//smallerOsm.Way = RelationsToWays(smallerOsm.Relation)
	//smallerOsm.Node = WaysToNodes(smallerOsm.Way)

	melbourne.BuildIndex()

	routes := nadeshiko.Routes{}
	routes.Activity("/", NewMainActivity())
	nadeshiko.Start(routes)
}
