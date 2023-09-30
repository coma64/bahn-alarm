import {Inject, Injectable} from '@angular/core';
import { Action, State, StateContext } from '@ngxs/store';
import { UserActions } from './user.actions';
import { Navigate } from '@ngxs/router-plugin';
import { User } from '../api';
import { EMPTY, Observable } from 'rxjs';
import { PushNotificationSubscriptionService } from '../modules/core/push-notification-subscription.service';
import LogRocket from "logrocket";
import {rollbarService} from "../rollbar";
import Rollbar from "rollbar";

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
    private readonly pushNotificationSubscription: PushNotificationSubscriptionService,
    @Inject(rollbarService) private readonly rollbar: Rollbar,
  ) {}

  @Action(UserActions.LoginSuccess)
  login(
    { patchState, dispatch }: StateContext<UserStateModel>,
    { user }: UserActions.LoginSuccess,
  ): Observable<unknown> {
    patchState({ user });
    dispatch(new Navigate(['/connections']));

    LogRocket.identify(user.id.toString(), {
      name: user.name,
    });

    this.rollbar.configure({payload: {user: {id: user.id, name: user.name}}})

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
