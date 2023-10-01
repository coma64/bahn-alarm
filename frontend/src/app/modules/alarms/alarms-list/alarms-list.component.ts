import { Component } from '@angular/core';
import { Select, Store } from '@ngxs/store';
import { Observable } from 'rxjs';
import { State } from '../../../state/state';
import { Urgency } from '../../../api';
import { AlarmsActions } from '../../../state/alarms.actions';
import { trackById } from '../../shared/track-by-id';
import { TimeSincePipe } from '../../shared/pipes/time-since.pipe';
import { SpinnerComponent } from '../../shared/components/spinner/spinner.component';
import { ConnectionInfoComponent } from '../connection-info/connection-info.component';
import { FeatherModule } from 'angular-feather';
import { NgIf, NgFor, AsyncPipe } from '@angular/common';

@Component({
    selector: 'app-alarms-list',
    templateUrl: './alarms-list.component.html',
    styleUrls: ['./alarms-list.component.scss'],
    standalone: true,
    imports: [
        NgIf,
        NgFor,
        FeatherModule,
        ConnectionInfoComponent,
        SpinnerComponent,
        AsyncPipe,
        TimeSincePipe,
    ],
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
