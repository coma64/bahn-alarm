import { Component, OnDestroy, OnInit } from '@angular/core';
import { TrackingService } from '../../../api';
import { Store } from '@ngxs/store';
import { ConnectionStats } from '../../../state/connection-stats.actions';
import { Connections } from '../../../state/connections.actions';
import { exhaustMap, Subject, takeUntil, timer } from 'rxjs';

@Component({
  selector: 'app-connections',
  templateUrl: './connections.component.html',
  styleUrls: ['./connections.component.scss'],
})
export class ConnectionsComponent implements OnInit, OnDestroy {
  private readonly destroy$ = new Subject<void>();

  constructor(
    private readonly tracking: TrackingService,
    private readonly store: Store,
  ) {}

  ngOnInit(): void {
    timer(0, 5_000)
      .pipe(
        exhaustMap(() => this.tracking.trackingStatsGet()),
        takeUntil(this.destroy$),
      )
      .subscribe((stats) =>
        this.store.dispatch(new ConnectionStats.Fetched(stats)),
      );

    timer(0, 5_000)
      .pipe(
        exhaustMap(() => this.tracking.trackingConnectionsGet()),
        takeUntil(this.destroy$),
      )
      .subscribe((response) =>
        this.store.dispatch(new Connections.Fetched(response.connections)),
      );
  }

  ngOnDestroy(): void {
    this.destroy$.next();
    this.destroy$.complete();
  }
}
