import {Injectable} from '@angular/core';
import {Recipe} from "./models/recipe.model";
import {HttpClient} from "@angular/common/http";

@Injectable({
  providedIn: 'root'
})
export class ApiService {
  private readonly url: string = "https://bakeoof.com/api";
  constructor(private http: HttpClient) {

  }

  /**
   * Get a recipe from the API server.
   */
  public async getRecipe(): Promise<Recipe> {
    let res: GoHttpResponse;
    try {
      res = <GoHttpResponse>(await this.http.get(this.url).toPromise());
    } catch (e) {
      throw new Error("Error getting results: " + e);
    }

    if (!res.ok) {
      throw new Error("Ok response not received: " + res.err);
    }

    return <Recipe>(res.data);
  }
}

class GoHttpResponse {
  public ok: boolean;
  public data?: any;
  public err?: string;
}
