import { Component, OnDestroy, OnInit } from '@angular/core';
import { TrackingService } from '../../../api';
import { Subject, takeUntil } from 'rxjs';
import { Store } from '@ngxs/store';
import { ConnectionStats } from '../../../state/connection-stats.actions';

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
    this.tracking
      .trackingStatsGet()
      .pipe(takeUntil(this.destroy$))
      .subscribe((stats) =>
        this.store.dispatch(new ConnectionStats.Fetched(stats)),
      );
  }

  ngOnDestroy(): void {
    this.destroy$.next();
    this.destroy$.complete();
  }
}
