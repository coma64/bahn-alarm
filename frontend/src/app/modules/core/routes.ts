import { Routes } from '@angular/router';
import { CoreComponent } from './core/core.component';

const routes: Routes = [
  {
    path: '',
    component: CoreComponent,
    children: [
      {
        path: 'connections',
        loadChildren: async () => import('../connections/routes'),
      },
      {
        path: 'profile',
        loadComponent: async () =>
          import('../profile/profile/profile.component'),
      },
      {
        path: 'alarms',
        loadComponent: async () => import('../alarms/alarms/alarms.component'),
      },
    ],
  },
  {
    path: '',
    pathMatch: 'full',
    redirectTo: '/connections',
  },
];

export default routes;
