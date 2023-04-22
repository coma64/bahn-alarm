import { inject } from '@angular/core';
import { CanActivateFn, Router } from '@angular/router';
import { Store } from '@ngxs/store';

export const isLoggedInGuard: CanActivateFn = () => {
  const isLoggedIn = !!inject(Store).selectSnapshot<string | undefined>(
    (state) => state.user.name,
  );

  if (isLoggedIn) return true;
  return inject(Router).parseUrl('/login');
};
