import { Injectable } from '@angular/core';
import { State, Action, StateContext } from '@ngxs/store';
import { User } from './user.actions';
import { Navigate } from '@ngxs/router-plugin';

export class UserStateModel {
  public readonly name?: string;
}

@State<UserStateModel>({
  name: 'user',
  defaults: {},
})
@Injectable()
export class UserState {
  @Action(User.LoginSuccess)
  login(
    { patchState, dispatch }: StateContext<UserStateModel>,
    { name }: User.LoginSuccess,
  ): void {
    patchState({ name });
    dispatch(new Navigate(['/connections']));
  }

  @Action(User.Logout)
  logout({ setState, dispatch }: StateContext<UserStateModel>): void {
    setState({});
    dispatch(new Navigate(['/login']));
  }
}
