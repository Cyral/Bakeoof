import {Injectable} from '@angular/core';
import {Recipe} from "./models/recipe.model";
import {HttpClient} from "@angular/common/http";

@Injectable({
  providedIn: 'root'
})
export class ApiService {
  private readonly url: string = "http://b2165353.ngrok.io";
  constructor(private http: HttpClient) {

  }

  /**
   * Get a recipe from the API server.
   */
  public async getRecipe(): Promise<Recipe> {
    const res = <GoHttpResponse>(await this.http.get(`${this.url}/test`).toPromise());
    if (!res.ok) {
      throw new Error("Ok response not received.");
    }
    return <Recipe>(res.data);

    /*
    return <Recipe>JSON.parse(
      `
      {"title":"This is a title","picture":"https://i.imgur.com/6ygY0ZK.png","story":"test test test this is a story about this","rating":5,"chef":"Boyardee","ingredients":["2 eggs","one cup of flour"],"steps":["Preheat the oven to 400","Burn the eggs","Pour flour on top"]}
      `
    )
    */
  }
}

class GoHttpResponse {
  public ok: boolean;
  public data: any;
}
