import { Injectable } from '@angular/core';
import { State, Action, StateContext } from '@ngxs/store';
import { AlarmsActions } from './alarms.actions';
import { Alarm, AlarmsService, Urgency } from '../api';
import { Observable, tap } from 'rxjs';

export class AlarmsStateModel {
  items: Array<Alarm> = [];
  page = 0;
  size = 50;
  filteredUrgency?: Urgency;
}

@State<AlarmsStateModel>({
  name: 'alarms',
})
@Injectable()
export class AlarmsState {
  constructor(private readonly alarms: AlarmsService) {}

  @Action(AlarmsActions.Fetch)
  fetch({
    getState,
    patchState,
  }: StateContext<AlarmsStateModel>): Observable<unknown> {
    const { page, size, filteredUrgency } = getState();

    return this.alarms.alarmsGet(page, size, filteredUrgency).pipe(
      tap({
        next: (response) => patchState({ items: response.alarms }),
        error: () =>
          alert('An unknown error occurred while loading the alarms'),
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
          patchState({ items: items.filter(({ id }) => id !== targetId) }),
        error: () =>
          alert('An unknown error occurred while deleting the alarm'),
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

  @Action(AlarmsActions.IncrementPage)
  incrementPage({
    getState,
    patchState,
    dispatch,
  }: StateContext<AlarmsStateModel>): Observable<any> {
    const { page } = getState();
    patchState({ page: page + 1 });

    return dispatch(AlarmsActions.Fetch);
  }

  @Action(AlarmsActions.DecrementPage)
  decrementPage({
    getState,
    patchState,
    dispatch,
  }: StateContext<AlarmsStateModel>): Observable<any> {
    const { page } = getState();
    patchState({ page: page - 1 });

    return dispatch(AlarmsActions.Fetch);
  }
}
