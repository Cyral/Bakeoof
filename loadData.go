package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// Recipee Definition of our struct that containts a recipee name,
// then a array of ingredients
type Recipee struct {
	Title       string   `{json:"name"}`
	Ingredients []string `{json:"ingredients"}`
}

// CreateRecipeeList returns the parsed json recipee list
func CreateRecipeeList() ([]Recipee, error) {
	jsonFile, err := os.Open("full_format_recipes.json")

	// check is see if we successfully loaded the json
	if err != nil {
		return nil, err
	}

	// Loads the json file into a byte file
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var recipees []Recipee

	// Loads the byte file into our recipee array
	json.Unmarshal(byteValue, &recipees)

	if err != nil {
		return nil, err
	}

	// Closes the json then returns the array
	defer jsonFile.Close()
	return recipees, nil
}
