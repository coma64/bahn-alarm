import { Component, OnInit } from '@angular/core';
import { AlarmsService } from '../../../api';
import { Store } from '@ngxs/store';
import { AlarmsActions } from '../../../state/alarms.actions';

@Component({
  selector: 'app-alarms',
  templateUrl: './alarms.component.html',
  styleUrls: ['./alarms.component.scss'],
})
export class AlarmsComponent implements OnInit {
  constructor(
    private readonly alarms: AlarmsService,
    private readonly store: Store,
  ) {}

  ngOnInit(): void {
    this.alarms.alarmsGet().subscribe({
      next: (response) =>
        this.store.dispatch(new AlarmsActions.Fetched(response.alarms)),
      error: () => alert('An unknown error occurred while loading the alarms'),
    });
  }
}
