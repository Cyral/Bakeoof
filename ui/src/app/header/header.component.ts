import {Component, OnInit} from '@angular/core';
import {ApiService} from "../api.service";
import {StateService} from "../state.service";

@Component({
  selector: 'app-header',
  templateUrl: "./header.component.html",
  styleUrls: ["./header.component.styl"]
})
export class HeaderComponent {

  public loading: boolean;

  constructor(private api: ApiService, private state: StateService) {

  }

  public async getRecipe() {
    if (this.loading) return;
    this.loading = true;

    const recipe = await this.api.getRecipe();
    this.state.recipeReceived.next(recipe);

    this.loading = false;
  }
}
