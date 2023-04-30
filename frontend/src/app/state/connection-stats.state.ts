import { Injectable } from '@angular/core';
import { State, Action, StateContext } from '@ngxs/store';
import { ConnectionStats } from './connection-stats.actions';
import { TrackingStats } from '../api';

export type TrackingStatsModel = TrackingStats | undefined;

@State<TrackingStatsModel>({
  name: 'connectionStats',
})
@Injectable()
export class ConnectionStatsState {
  @Action(ConnectionStats.Fetched)
  add(
    { setState }: StateContext<TrackingStatsModel>,
    stats: ConnectionStats.Fetched,
  ) {
    setState(stats.stats);
  }
}
