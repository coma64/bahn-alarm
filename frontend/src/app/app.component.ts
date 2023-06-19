import { Component, OnDestroy, OnInit } from '@angular/core';
import { SwPush, SwUpdate } from '@angular/service-worker';
import { exhaustMap, filter, interval, Subject, takeUntil } from 'rxjs';
import { fromPromise } from 'rxjs/internal/observable/innerFrom';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss'],
})
export class AppComponent implements OnInit, OnDestroy {
  private readonly destroy$ = new Subject<void>();

  constructor(private readonly swUpdate: SwUpdate) {}

  ngOnInit(): void {
    // if a new version is available it is handled by the below subscription
    interval(5 * 60_000)
      .pipe(
        exhaustMap(() => fromPromise(this.swUpdate.checkForUpdate())),
        takeUntil(this.destroy$),
      )
      .subscribe();

    this.swUpdate.versionUpdates
      .pipe(
        filter((e) => e.type === 'VERSION_READY'),
        takeUntil(this.destroy$),
      )
      .subscribe(
        () =>
          confirm(
            'A new version of this app is available. Do you want to reload?',
          ) && location.reload(),
      );
  }

  ngOnDestroy(): void {
    this.destroy$.next();
    this.destroy$.complete();
  }
}
