import { Injectable } from '@angular/core';
import { State, Action, StateContext } from '@ngxs/store';
import { ConnectionStats } from './connection-stats.actions';
import { TrackingStats } from '../api';

export type ConnectionStatsModel = {
  stats?: TrackingStats;
};

@State<ConnectionStatsModel>({
  name: 'connectionStats',
})
@Injectable()
export class ConnectionStatsState {
  @Action(ConnectionStats.Fetched)
  add(
    { setState }: StateContext<ConnectionStatsModel>,
    stats: ConnectionStats.Fetched,
  ) {
    setState({ stats: stats.stats });
  }
}
