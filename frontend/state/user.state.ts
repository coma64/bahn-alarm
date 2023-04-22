import { Injectable } from '@angular/core';
import { State, Action, StateContext } from '@ngxs/store';
import { LoginSuccess } from './user.actions';

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
  add({ patchState }: StateContext<UserStateModel>, { name }: LoginSuccess) {
    patchState({ name });
  }
}
