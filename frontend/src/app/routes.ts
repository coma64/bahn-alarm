import { Routes } from '@angular/router';
import { isLoggedInGuard } from './guards/logged-in.guard';

export const routes: Routes = [
  {
    path: '',
    pathMatch: 'full',
    redirectTo: 'connections',
  },
  {
    path: 'login',
    loadComponent: async () => import('./modules/login/login/login.component'),
  },
  {
    path: '',
    loadChildren: async () => import('./modules/core/routes'),
    canActivate: [isLoggedInGuard],
  },
  {
    path: '**',
    redirectTo: 'connections',
  },
];
