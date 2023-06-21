import { Component, OnDestroy, OnInit } from '@angular/core';
import { Select, Store } from '@ngxs/store';
import { AlarmsActions } from '../../../state/alarms.actions';
import { AlarmsState, AlarmsStateModel } from '../../../state/alarms.state';
import { exhaustMap, Observable, Subject, takeUntil, timer } from 'rxjs';

@Component({
  selector: 'app-alarms',
  templateUrl: './alarms.component.html',
  styleUrls: ['./alarms.component.scss'],
})
export class AlarmsComponent implements OnInit, OnDestroy {
  @Select(AlarmsState)
  protected readonly alarms$!: Observable<AlarmsStateModel>;

  private readonly destroy$ = new Subject<void>();

  constructor(private readonly store: Store) {}

  ngOnInit() {
    timer(0, 5_000)
      .pipe(
        exhaustMap(() => this.store.dispatch(AlarmsActions.Fetch)),
        takeUntil(this.destroy$),
      )
      .subscribe();
  }

  ngOnDestroy(): void {
    this.destroy$.next();
    this.destroy$.complete();
  }

  changePage(newPage: number): void {
    this.store.dispatch(new AlarmsActions.SetPage(newPage));
  }
}
