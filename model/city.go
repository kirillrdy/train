package model

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type City struct {
	TrainLines TrainLines
}

func LoadCity(name string) (city City) {
	file, err := os.Open(name)
	if err != nil {
		log.Fatalf("ERROR: %v \n", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	decoder.Decode(&city)

	return city
}

func (city City) Save(name string) {
	json_file, err := os.Create(name)
	if err != nil {
		panic(err)
	}

	jsonEncoder := json.NewEncoder(json_file)
	jsonEncoder.Encode(city)
	json_file.Close()

	IndentJson(name)
}

func LoadJson(name string) (data []byte) {
	file, err := os.Open(name)
	defer file.Close()
	if err != nil {
		log.Fatalf("ERROR %v \n", err)
	}
	data, err = ioutil.ReadAll(file)

	if err != nil {
		log.Fatalf("ERROR %v \n", err)
	}

	return data
}

func IndentJson(name string) {
	data := LoadJson(name)
	var buffer bytes.Buffer
	json.Indent(&buffer, data, "", "  ")

	file, err := os.Create(name)
	if err != nil {
		log.Fatalf("ERROR %v \n", err)
	}
	defer file.Close()

	file.WriteString(buffer.String())

}
