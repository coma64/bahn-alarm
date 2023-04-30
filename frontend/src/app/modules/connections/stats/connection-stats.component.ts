import { Component } from '@angular/core';
import { Observable } from 'rxjs';
import {
  ConnectionStatsState,
  ConnectionStatsModel,
} from '../../../state/connection-stats.state';
import { Select } from '@ngxs/store';

@Component({
  selector: 'app-connection-stats',
  templateUrl: './connection-stats.component.html',
  styleUrls: ['./connection-stats.component.scss'],
})
export class ConnectionStatsComponent {
  @Select(ConnectionStatsState)
  readonly stats$!: Observable<ConnectionStatsModel>;
}
