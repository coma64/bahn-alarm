import { Component, OnInit } from '@angular/core';
import { TrackingService } from '../../../api';
import { Store } from '@ngxs/store';
import { ConnectionStats } from '../../../state/connection-stats.actions';
import { Connections } from '../../../state/connections.actions';
import { State } from '../../../state/state';
import {
  ConnectionStatsModel,
  ConnectionStatsState,
} from '../../../state/connection-stats.state';

@Component({
  selector: 'app-connections',
  templateUrl: './connections.component.html',
  styleUrls: ['./connections.component.scss'],
})
export class ConnectionsComponent implements OnInit {
  constructor(
    private readonly tracking: TrackingService,
    private readonly store: Store,
  ) {}

  ngOnInit(): void {
    if (
      !this.store.selectSnapshot<ConnectionStatsModel>(ConnectionStatsState)
        .stats
    )
      this.tracking
        .trackingStatsGet()
        .subscribe((stats) =>
          this.store.dispatch(new ConnectionStats.Fetched(stats)),
        );

    if (this.store.selectSnapshot((s: State) => s.connections.page) === -1)
      this.tracking
        .trackingConnectionsGet()
        .subscribe((response) =>
          this.store.dispatch(new Connections.Fetched(response.connections)),
        );
  }
}
