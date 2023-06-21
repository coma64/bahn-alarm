import { Injectable } from '@angular/core';
import { State, Action, StateContext } from '@ngxs/store';
import { AlarmsActions } from './alarms.actions';
import { Alarm, AlarmsService, Urgency } from '../api';
import { EMPTY, Observable, tap } from 'rxjs';
import { NotifyService } from '../modules/shared/services/notify.service';

export class AlarmsStateModel {
  items?: Alarm[];
  page = 0;
  maxPage = 0;
  size = 50;
  filteredUrgency?: Urgency;
}

@State<AlarmsStateModel>({
  name: 'alarms',
  defaults: new AlarmsStateModel(),
})
@Injectable()
export class AlarmsState {
  constructor(
    private readonly alarms: AlarmsService,
    private readonly notify: NotifyService,
  ) {}

  @Action(AlarmsActions.Fetch)
  fetch({
    getState,
    patchState,
  }: StateContext<AlarmsStateModel>): Observable<unknown> {
    const { page, size, filteredUrgency } = getState();

    return this.alarms.alarmsGet(page, size, filteredUrgency).pipe(
      tap({
        next: (response) =>
          patchState({
            items: response.alarms,
            maxPage: Math.ceil(response.pagination.totalItems / size) - 1,
          }),
        error: () =>
          this.notify.error(
            'An unknown error occurred while loading the alarms',
          ),
      }),
    );
  }

  @Action(AlarmsActions.Delete)
  delete(
    { getState, patchState }: StateContext<AlarmsStateModel>,
    { targetId }: AlarmsActions.Delete,
  ): Observable<unknown> {
    const { items } = getState();
    return this.alarms.alarmsIdDelete(targetId).pipe(
      tap({
        next: () =>
          patchState({ items: items?.filter(({ id }) => id !== targetId) }),
        error: () =>
          this.notify.error(
            'An unknown error occurred while deleting the alarm',
          ),
      }),
    );
  }

  @Action(AlarmsActions.FilterByUrgency)
  filterByUrgency(
    { patchState, dispatch }: StateContext<AlarmsStateModel>,
    { urgency }: AlarmsActions.FilterByUrgency,
  ): Observable<any> {
    patchState({ filteredUrgency: urgency });
    return dispatch(AlarmsActions.Fetch);
  }

  @Action(AlarmsActions.SetPage)
  setPage(
    { getState, patchState, dispatch }: StateContext<AlarmsStateModel>,
    { newPage }: AlarmsActions.SetPage,
  ): Observable<any> {
    const { page, maxPage } = getState();
    if (newPage > maxPage) newPage = maxPage;
    if (newPage < 0) newPage = 0;
    if (newPage === page) return EMPTY;

    patchState({ page: newPage });
    return dispatch(AlarmsActions.Fetch);
  }
}
