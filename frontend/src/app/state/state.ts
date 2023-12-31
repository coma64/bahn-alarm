import { UserStateModel } from './user.state';
import { ConnectionStatsModel } from './connection-stats.state';
import { ConnectionsStateModel } from './connections.state';
import { AlarmedDevicesStateModel } from './alarmed-devices.state';
import { AlarmsStateModel } from './alarms.state';

export type State = {
  user: UserStateModel;
  connectionStats: ConnectionStatsModel;
  connections: ConnectionsStateModel;
  alarmedDevices: AlarmedDevicesStateModel;
  alarms: AlarmsStateModel;
};
