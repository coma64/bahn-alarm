import { Injectable } from '@angular/core';
import { State, Action, StateContext } from '@ngxs/store';
import { LoginSuccess } from './user.actions';
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
  @Action(LoginSuccess)
  add(
    { patchState, dispatch }: StateContext<UserStateModel>,
    { name }: LoginSuccess,
  ) {
    patchState({ name });
    if (name) dispatch(new Navigate(['/connections']));
    else dispatch(new Navigate(['/login']));
  }
}
