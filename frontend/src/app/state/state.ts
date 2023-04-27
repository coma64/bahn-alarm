import { UserStateModel } from './user.state';
import { TrackingStatsModel } from './connection-stats.state';
import { ConnectionsStateModel } from './connections.state';
import { AlarmedDevicesStateModel } from './alarmed-devices.state';

export interface State {
  user: UserStateModel;
  connectionStats: TrackingStatsModel;
  connections: ConnectionsStateModel;
  alarmedDevices: AlarmedDevicesStateModel;
}
