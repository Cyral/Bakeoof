import {Component, Input} from '@angular/core';

@Component({
  selector: 'app-button',
  template: `
    <div>
      <ng-content></ng-content>
      <span class="loader" *ngIf="loading">
        <i class="fa fa-circle-o-notch fa-spin" aria-hidden="true"></i>
      </span>
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
