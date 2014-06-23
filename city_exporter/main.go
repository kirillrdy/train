package main

import (
	"github.com/kirillrdy/osm"
	"github.com/kirillrdy/train/model"
)

func main() {
	melbourneOsm := osm.LoadPackagedMelbourne()
	melbourneOsm.BuildIndex()
	melbourne := model.City{}

	trainLines := model.AllTrainLines(*melbourneOsm)
	melbourne.TrainLines = trainLines
	melbourne.Save("melbourne.json")
}
