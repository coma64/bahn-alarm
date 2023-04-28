import { Injectable } from '@angular/core';
import { State, Action, StateContext } from '@ngxs/store';
import { AlarmsActions } from './alarms.actions';
import { Alarm } from '../api';

export interface AlarmsStateModel {
  items: Array<Alarm>;
  page: number;
}

const defaults = {
  items: [],
  page: -1,
};

@State<AlarmsStateModel>({
  name: 'alarms',
  defaults,
})
@Injectable()
export class AlarmsState {
  @Action(AlarmsActions.Fetched)
  fetched(
    { getState, setState }: StateContext<AlarmsStateModel>,
    { items }: AlarmsActions.Fetched,
  ): void {
    const { page } = getState();
    setState({ items, page });
  }

  @Action(AlarmsActions.Deleted)
  deleted(
    { getState, setState }: StateContext<AlarmsStateModel>,
    { deletedId }: AlarmsActions.Deleted,
  ): void {
    const { items, page } = getState();
    setState({ items: items.filter(({ id }) => id !== deletedId), page });
  }
}
