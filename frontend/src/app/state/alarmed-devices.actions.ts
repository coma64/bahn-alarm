import { PushNotificationSubscription } from '../api';

export namespace AlarmedDeviceActions {
  export class Fetched {
    static readonly type = '[AlarmedDevices] Fetched';
    constructor(readonly devices: PushNotificationSubscription[]) {}
  }
}
