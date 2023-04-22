import { Component } from '@angular/core';
import { TrackedConnection } from '../../../api';
import { State } from '../../../state/state';
import { Select } from '@ngxs/store';
import { Observable } from 'rxjs';

@Component({
  selector: 'app-connection-list',
  templateUrl: './connection-list.component.html',
  styleUrls: ['./connection-list.component.scss'],
})
export class ConnectionListComponent {
  @Select((state: State) => state.connections.items)
  readonly connections$!: Observable<Array<TrackedConnection>>;
}
