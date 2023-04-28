import { Component } from '@angular/core';
import { Select, Store } from '@ngxs/store';
import { Observable } from 'rxjs';
import { State } from '../../../state/state';
import { AlarmsService, Urgency } from '../../../api';
import { AlarmsActions } from '../../../state/alarms.actions';

@Component({
  selector: 'app-alarms-list',
  templateUrl: './alarms-list.component.html',
  styleUrls: ['./alarms-list.component.scss'],
})
export class AlarmsListComponent {
  @Select() readonly alarms$!: Observable<State['alarms']>;
  readonly emergencyToHumanReadable: Record<Urgency, string> = {
    info: 'Info',
    warn: 'Warning',
    error: 'Error',
  };

  constructor(
    private readonly alarms: AlarmsService,
    private readonly store: Store,
  ) {}

  delete(id: number): void {
    this.alarms.alarmsIdDelete(id).subscribe({
      next: () => this.store.dispatch(new AlarmsActions.Deleted(id)),
      error: () => alert('An unknown error occurred while delete the alarm'),
    });
  }
}
