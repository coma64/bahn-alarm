import { Component } from '@angular/core';
import { Observable } from 'rxjs';
import {
  ConnectionStatsState,
  ConnectionStatsModel,
} from '../../../state/connection-stats.state';
import { Select } from '@ngxs/store';
import { NextRelativeTimePipe } from '../../shared/pipes/next-relative-time.pipe';
import { ToRelativeTimePipe } from '../../shared/pipes/to-relative-time.pipe';
import { SpinnerComponent } from '../../shared/components/spinner/spinner.component';
import { NgIf, AsyncPipe } from '@angular/common';

@Component({
    selector: 'app-connection-stats',
    templateUrl: './connection-stats.component.html',
    styleUrls: ['./connection-stats.component.scss'],
    standalone: true,
    imports: [
        NgIf,
        SpinnerComponent,
        AsyncPipe,
        ToRelativeTimePipe,
        NextRelativeTimePipe,
    ],
})
export class ConnectionStatsComponent {
  @Select(ConnectionStatsState)
  readonly stats$!: Observable<ConnectionStatsModel>;
}
