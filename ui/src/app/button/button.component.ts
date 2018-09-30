import {Component, Input} from '@angular/core';

@Component({
  selector: 'app-button',
  template: `
    <div>
      <span class="loader" *ngIf="loading">
        <i class="fa fa-circle-o-notch fa-spin" aria-hidden="true"></i>
      </span>
      <ng-content></ng-content>
    </div>
  `,
  styleUrls: ['./button.component.styl']
})
export class ButtonComponent {
  @Input()
  public loading?: boolean;

  constructor() {

  }
}
