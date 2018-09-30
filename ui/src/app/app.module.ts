import {BrowserModule} from '@angular/platform-browser';
import {NgModule} from '@angular/core';

import {AppComponent} from './app.component';
import {HeaderComponent} from './header/header.component';
import {ButtonComponent} from "./button/button.component";
import {ApiService} from "./api.service";
import { RecipeComponent } from './recipe/recipe.component';
import {HttpClientModule} from "@angular/common/http";

@NgModule({
  declarations: [
    AppComponent,
    HeaderComponent,
    ButtonComponent,
    RecipeComponent
  ],
  imports: [
    BrowserModule,
    HttpClientModule,
  ],
  providers: [
    ApiService
  ],
  bootstrap: [AppComponent]
})
export class AppModule {
}
