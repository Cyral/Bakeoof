package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

// Recipee Definition of our struct that containts a recipee name,
// then a array of ingredients
type Recipe struct {
	Title       string   `{json:"name"}`
	Ingredients []string `{json:"ingredients"}`
	BakeType    string
}

// CreateRecipeeList returns the parsed json recipee list
func CreateRecipeList() ([]Recipe, error) {
	jsonFile, err := os.Open("full_format_recipes.json")

	// check is see if we successfully loaded the json
	if err != nil {
		return nil, err
	}

	// Loads the json file into a byte file
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var recipes []Recipe

	// Loads the byte file into our recipee array
	json.Unmarshal(byteValue, &recipes)

	if err != nil {
		return nil, err
	}

	// Closes the json then returns the array
	defer jsonFile.Close()
	return recipes, nil
}

func GetRecipe() (Recipe, error) {

	recipes, err := CreateRecipeList()
	var numRecipes int = len(recipes)

	// Initializes our random number generator
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	// Instantiating the new recipee we will be creating
	var newRecipe Recipe
	var finalRecipeName bytes.Buffer

	// Determines the number of ingredients our new recipe will have
	numIngredients := r1.Intn(4) + 3

	// The different types of baked goods we will randomly choose from
	bakeTypes := [3]string{"Cake", "Muffins", "Cookies"}
	var bakedChoice int = r1.Intn(3)

	// Denotes the recipee we will iteratively randomly choose from our list of recipes
	// as well as the number of ingredients that recipe has (so we can choose one from them)
	var choiceIngredient int
	var choiceRecipe int

	// Resize the array of ingredients for our return recipe
	var newIngredientList = make([]string, numIngredients)
	newRecipe.Ingredients = newIngredientList

	for i := 0; i < numIngredients; i++ {
		choiceRecipe = r1.Intn(numRecipes)
		choiceIngredient = (r1.Intn(100) + 1) % len(recipes[choiceRecipe].Ingredients)

		newRecipe.Ingredients[i] = recipes[choiceRecipe].Ingredients[choiceIngredient]

		// Our final recipe will share the name our recipee from which we pull the first ingredient
		if i == 0 {
			finalRecipeName.WriteString(recipes[choiceRecipe].Title)
		}
	}

	if err != nil {
		return newRecipe, err
	}

	finalRecipeName.WriteString(bakeTypes[bakedChoice])
	newRecipe.Title = finalRecipeName.String()
	newRecipe.BakeType = bakeTypes[bakedChoice]

	return newRecipe, nil
}
