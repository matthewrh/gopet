package utils

import (
	"encoding/json"
	"github.com/matthewrh/gopet/pkg/pet"
	"io/ioutil"
)

func LoadPet(filename string) (*pet.Pet, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var pet pet.Pet
	if err := json.Unmarshal(data, &pet); err != nil {
		return nil, err
	}
	return &pet, nil
}
