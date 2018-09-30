import {Injectable} from '@angular/core';
import {Recipe} from "./models/recipe.model";
import {Subject} from "rxjs";

@Injectable({
  providedIn: 'root'
})
export class StateService {

  public recipeReceived: Subject<Recipe>;

  constructor() {
    this.recipeReceived = new Subject<Recipe>();
  }
}
