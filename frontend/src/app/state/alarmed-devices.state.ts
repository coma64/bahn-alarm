import { Injectable } from '@angular/core';
import { State, Action, StateContext } from '@ngxs/store';
import { AlarmedDeviceActions } from './alarmed-devices.actions';
import { PushNotificationSubscription } from '../api';

export interface AlarmedDevicesStateModel {
  items: Array<PushNotificationSubscription>;
  page: number;
}

const defaults = {
  items: [],
  page: -1,
};

@State<AlarmedDevicesStateModel>({
  name: 'alarmedDevices',
  defaults,
})
@Injectable()
export class AlarmedDevicesState {
  @Action(AlarmedDeviceActions.Fetched)
  fetched(
    { getState, setState }: StateContext<AlarmedDevicesStateModel>,
    { devices }: AlarmedDeviceActions.Fetched,
  ) {
    const { page } = getState();
    setState({ items: devices, page: page + 1 });
  }
}
