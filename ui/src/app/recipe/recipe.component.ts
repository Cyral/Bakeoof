import {Component, OnInit} from '@angular/core';
import {StateService} from "../state.service";
import {Recipe} from "../models/recipe.model";

@Component({
  selector: 'app-recipe',
  templateUrl: "./recipe.component.html",
  styleUrls: ["./recipe.component.styl"]
})
export class RecipeComponent {

  public recipe: Recipe;

  constructor(private state: StateService) {
    state.recipeReceived.subscribe(this.onRecipeReceived.bind(this));
  }

  public getStars() {
    return Array(Math.round(this.recipe.rating)).fill(1);
  }

  private onRecipeReceived(recipe: Recipe) {
    this.recipe = recipe;
  }
}
