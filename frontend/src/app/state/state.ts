import { UserStateModel } from './user.state';
import { TrackingStatsModel } from './connection-stats.state';

export interface State {
  user: UserStateModel;
  connectionStats: TrackingStatsModel;
}
