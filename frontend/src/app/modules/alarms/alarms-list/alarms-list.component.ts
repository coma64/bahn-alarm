import { Component } from '@angular/core';
import { Store } from '@ngxs/store';
import { Urgency } from '../../../api';
import { AlarmsActions } from '../../../state/alarms.actions';
import { trackById } from '../../shared/track-by-id';
import { TimeSincePipe } from '../../shared/pipes/time-since.pipe';
import { SpinnerComponent } from '../../shared/components/spinner/spinner.component';
import { ConnectionInfoComponent } from '../connection-info/connection-info.component';
import { AsyncPipe, NgFor, NgIf } from '@angular/common';
import { IconsModule } from '../../icons/icons.module';
import { toSignal } from '@angular/core/rxjs-interop';
import { alarmsStateToken } from '../../../state/alarms.state';

@Component({
  selector: 'app-alarms-list',
  templateUrl: './alarms-list.component.html',
  styleUrls: ['./alarms-list.component.scss'],
  standalone: true,
  imports: [
    NgIf,
    NgFor,
    IconsModule,
    ConnectionInfoComponent,
    SpinnerComponent,
    AsyncPipe,
    TimeSincePipe,
  ],
})
export class AlarmsListComponent {
  protected readonly alarms = toSignal(this.store.select(alarmsStateToken));

  protected readonly emergencyToHumanReadable: Record<Urgency, string> = {
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
