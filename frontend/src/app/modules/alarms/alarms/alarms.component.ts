import { Component, OnDestroy, OnInit } from '@angular/core';
import { Select, Store } from '@ngxs/store';
import { AlarmsActions } from '../../../state/alarms.actions';
import { AlarmsState, AlarmsStateModel } from '../../../state/alarms.state';
import { exhaustMap, Observable, Subject, timer } from 'rxjs';
import { PaginationComponent } from '../../shared/components/pagination/pagination.component';
import { AsyncPipe, NgIf } from '@angular/common';
import { AlarmsListComponent } from '../alarms-list/alarms-list.component';
import { FiltersComponent } from '../filters/filters.component';
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';

@Component({
  selector: 'app-alarms',
  templateUrl: './alarms.component.html',
  styleUrls: ['./alarms.component.scss'],
  standalone: true,
  imports: [
    FiltersComponent,
    AlarmsListComponent,
    NgIf,
    PaginationComponent,
    AsyncPipe,
  ],
})
export default class AlarmsComponent implements OnInit, OnDestroy {
  @Select(AlarmsState)
  protected readonly alarms$!: Observable<AlarmsStateModel>;

  private readonly destroy$ = new Subject<void>();
  private readonly untilDestroyed = takeUntilDestroyed();

  constructor(private readonly store: Store) {}

  ngOnInit() {
    timer(0, 5_000)
      .pipe(
        exhaustMap(() => this.store.dispatch(AlarmsActions.Fetch)),
        this.untilDestroyed,
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
