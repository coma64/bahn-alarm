import { inject } from '@angular/core';
import { CanActivateFn, Router } from '@angular/router';
import { Store } from '@ngxs/store';
import { State } from '../state/state';

export const isLoggedInGuard: CanActivateFn = () => {
  const snapshot: State = inject(Store).snapshot();

  if (snapshot.user) return true;
  return inject(Router).parseUrl('/login');
};
