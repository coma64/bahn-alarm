import { Injectable } from '@angular/core';
import { State, Action, StateContext } from '@ngxs/store';
import { Connections } from './connections.actions';
import { TrackedConnection } from '../api';

export type ConnectionsStateModel = {
  items: TrackedConnection[];
  page: number;
};

const defaults = {
  items: [],
  page: -1,
};

@State<ConnectionsStateModel>({
  name: 'connections',
  defaults,
})
@Injectable()
export class ConnectionsState {
  @Action(Connections.Fetched)
  fetched(
    { getState, setState }: StateContext<ConnectionsStateModel>,
    { connections }: Connections.Fetched,
  ): void {
    setState({ items: connections, page: getState().page + 1 });
  }

  @Action(Connections.Created)
  created(
    { getState, setState }: StateContext<ConnectionsStateModel>,
    { newConnection }: Connections.Created,
  ): void {
    const { items, page } = getState();
    setState({ items: [...items, newConnection], page });
  }

  @Action(Connections.Updated)
  updated(
    { getState, patchState }: StateContext<ConnectionsStateModel>,
    { updatedConnection }: Connections.Updated,
  ): void {
    patchState({
      items: getState().items.map((c) =>
        c.id === updatedConnection.id ? updatedConnection : c,
      ),
    });
  }
}
