import { Injectable } from '@angular/core';
import { State, Action, StateContext } from '@ngxs/store';
import { UserActions } from './user.actions';
import { Navigate } from '@ngxs/router-plugin';
import { User } from '../api';
import { patch } from '@ngxs/store/operators';

export type UserStateModel = {
  user?: User;
  hasDeniedPushNotifications: boolean;
};

@State<UserStateModel>({
  name: 'user',
  defaults: { hasDeniedPushNotifications: false },
})
@Injectable()
export class UserState {
  @Action(UserActions.LoginSuccess)
  login(
    { patchState, dispatch }: StateContext<UserStateModel>,
    { user }: UserActions.LoginSuccess,
  ): void {
    patchState({ user });
    dispatch(new Navigate(['/connections']));
  }

  @Action(UserActions.Logout)
  logout({ patchState, dispatch }: StateContext<UserStateModel>): void {
    patchState({ user: undefined });
    dispatch(new Navigate(['/login']));
  }

  @Action(UserActions.DeniedPushNotification)
  deniedPushNotification({ patchState }: StateContext<UserStateModel>): void {
    patchState({ hasDeniedPushNotifications: true });
  }
}
