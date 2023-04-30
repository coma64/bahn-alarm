import { Component, OnInit } from '@angular/core';
import { Store } from '@ngxs/store';
import { AlarmsActions } from '../../../state/alarms.actions';

@Component({
  selector: 'app-alarms',
  templateUrl: './alarms.component.html',
  styleUrls: ['./alarms.component.scss'],
})
export class AlarmsComponent implements OnInit {
  constructor(private readonly store: Store) {}

  ngOnInit() {
    this.store.dispatch(AlarmsActions.Fetch);
  }
}
