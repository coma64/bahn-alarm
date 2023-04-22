import { Injectable } from '@angular/core';
import { State, Action, StateContext } from '@ngxs/store';
import { Connections } from './connections.actions';
import { TrackedConnection } from '../../api';

export class ConnectionsStateModel {
  public items: TrackedConnection[] = [];
  public page = 0;
}

const defaults = {
  items: [],
  page: 0,
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
  ) {
    setState({ items: connections, page: getState().page + 1 });
  }
}
