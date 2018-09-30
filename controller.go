package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
)

type okResponse struct {
	Ok   bool        `json:"ok"`
	Data interface{} `json:"data"`
}

type errResponse struct {
	Ok    bool   `json:"ok"`
	Error string `json:"error"`
}

// RecipeOutput is what is sent to the front end
type RecipeOutput struct {
	Title       string   `json:"title"`
	Picture     string   `json:"picture"`
	Story       string   `json:"story"`
	Rating      int      `json:"rating"`
	Chef        string   `json:"chef"`
	Ingredients []string `json:"ingredients"`
	Steps       []string `json:"steps"`
}

// RootController handles the / endpoint
func RootController(w http.ResponseWriter, r *http.Request) {
	sendResponse(getRecipe(w), w)
}

func getRecipe(w http.ResponseWriter) RecipeOutput {
	recipe, err := GetRecipe()
	if err != nil {
		sendError(err, w)
		return RecipeOutput{}
	}

	return RecipeOutput{
		Title:       recipe.Title,
		Picture:     GetPictureURL(recipe),
		Story:       "",
		Rating:      rand.Intn(5) + 1,
		Chef:        GetChef(),
		Ingredients: recipe.Ingredients,
		Steps:       GetSteps(recipe),
	}
}

func sendResponse(data interface{}, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	json.NewEncoder(w).Encode(okResponse{Ok: true, Data: data})
}

func sendError(err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(errResponse{Ok: false, Error: err.Error()})
}
