package model

import (
	"encoding/json"
	"os"
)

type City struct {
	TrainLines TrainLines
}

func LoadCity(name string) (city City) {
	file, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	decoder.Decode(&city)

	return city
}

func (city City) Save(name string) {
	json_file, err := os.Create(name)
	defer json_file.Close()
	if err != nil {
		panic(err)
	}

	jsonEncoder := json.NewEncoder(json_file)
	jsonEncoder.Encode(city)
}
