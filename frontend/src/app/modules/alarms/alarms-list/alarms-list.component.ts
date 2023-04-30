import { Component } from '@angular/core';
import { Select, Store } from '@ngxs/store';
import { Observable } from 'rxjs';
import { State } from '../../../state/state';
import { Urgency } from '../../../api';
import { AlarmsActions } from '../../../state/alarms.actions';
import { trackById } from '../../shared/track-by-id';

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

  protected readonly trackById = trackById;

  constructor(private readonly store: Store) {}

  delete(id: number): void {
    this.store.dispatch(new AlarmsActions.Delete(id));
  }
}
