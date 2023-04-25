import { UserStateModel } from './user.state';
import { TrackingStatsModel } from './connection-stats.state';
import { ConnectionsStateModel } from './connections.state';

export interface State {
  user: UserStateModel;
  connectionStats: TrackingStatsModel;
  connections: ConnectionsStateModel;
}
