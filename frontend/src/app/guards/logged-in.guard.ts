import { inject } from '@angular/core';
import { CanActivateFn, Router } from '@angular/router';
import { Store } from '@ngxs/store';
import { UserState, UserStateModel } from '../state/user.state';

export const isLoggedInGuard: CanActivateFn = () => {
  const user = inject(Store).selectSnapshot<UserStateModel>(UserState);

  if (user) return true;
  return inject(Router).parseUrl('/login');
};
