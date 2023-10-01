import { Component, TrackByFunction } from '@angular/core';
import { TrackedConnection, TrackedDeparture } from '../../../api';
import { State } from '../../../state/state';
import { Select } from '@ngxs/store';
import { Observable } from 'rxjs';
import { Router } from '@angular/router';
import { trackById } from '../../shared/track-by-id';
import { SpinnerComponent } from '../../shared/components/spinner/spinner.component';
import { DepartureComponent } from '../departure/departure.component';
import { NgIf, NgFor, AsyncPipe } from '@angular/common';
import { FeatherModule } from 'angular-feather';

@Component({
    selector: 'app-connection-list',
    templateUrl: './connection-list.component.html',
    styleUrls: ['./connection-list.component.scss'],
    standalone: true,
    imports: [
        FeatherModule,
        NgIf,
        NgFor,
        DepartureComponent,
        SpinnerComponent,
        AsyncPipe,
    ],
})
export class ConnectionListComponent {
  @Select((state: State) => state.connections.items)
  protected readonly connections$!: Observable<TrackedConnection[]>;

  protected readonly trackById = trackById;

  constructor(public readonly router: Router) {}

  trackByDeparture: TrackByFunction<TrackedDeparture> = (_, { departure }) =>
    departure;
}
