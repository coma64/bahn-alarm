import { Injectable } from '@angular/core';
import { State, Action, StateContext } from '@ngxs/store';
import { UserActions } from './user.actions';
import { Navigate } from '@ngxs/router-plugin';
import { User } from '../api';

export type UserStateModel = User | undefined;

@State<UserStateModel>({
  name: 'user',
})
@Injectable()
export class UserState {
  @Action(UserActions.LoginSuccess)
  login(
    { setState, dispatch }: StateContext<UserStateModel>,
    { user }: UserActions.LoginSuccess,
  ): void {
    setState(user);
    dispatch(new Navigate(['/connections']));
  }

  @Action(UserActions.Logout)
  logout({ setState, dispatch }: StateContext<UserStateModel>): void {
    setState(undefined);
    dispatch(new Navigate(['/login']));
  }
}
