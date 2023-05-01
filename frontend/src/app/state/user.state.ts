import { Injectable } from '@angular/core';
import { Action, State, StateContext } from '@ngxs/store';
import { UserActions } from './user.actions';
import { Navigate } from '@ngxs/router-plugin';
import { User } from '../api';
import { EMPTY, Observable } from 'rxjs';
import { NotifyService } from '../modules/shared/services/notify.service';
import { PushNotificationSubscriptionService } from '../modules/core/push-notification-subscription.service';

export type UserStateModel = {
  user?: User;
  pushSubId?: number;
};

@State<UserStateModel>({
  name: 'user',
})
@Injectable()
export class UserState {
  constructor(
    private readonly notify: NotifyService,
    private readonly pushNotificationSubscription: PushNotificationSubscriptionService,
  ) {}

  @Action(UserActions.LoginSuccess)
  login(
    { patchState, dispatch }: StateContext<UserStateModel>,
    { user }: UserActions.LoginSuccess,
  ): Observable<unknown> {
    patchState({ user });
    dispatch(new Navigate(['/connections']));

    return this.pushNotificationSubscription.askUserAndRegister();
  }

  @Action(UserActions.Logout)
  logout({
    setState,
    getState,
    dispatch,
  }: StateContext<UserStateModel>): Observable<unknown> {
    const { pushSubId } = getState();
    setState({});
    dispatch(new Navigate(['/login']));

    if (!pushSubId) return EMPTY;

    return this.pushNotificationSubscription.unregister(pushSubId);
  }

  @Action(UserActions.RegisteredPushNotifications)
  registeredPushNotifications(
    { patchState }: StateContext<UserStateModel>,
    payload: UserActions.RegisteredPushNotifications,
  ): void {
    patchState({
      pushSubId: payload.pushSubId,
    });
  }
}
