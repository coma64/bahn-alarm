import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { CoreComponent } from './core/core.component';

const routes: Routes = [
  {
    path: '',
    component: CoreComponent,
    children: [
      {
        path: 'connections',
        loadChildren: () =>
          import('../connections/connections.module').then(
            (m) => m.ConnectionsModule,
          ),
      },
      {
        path: 'profile',
        loadChildren: () =>
          import('../profile/profile.module').then((m) => m.ProfileModule),
      },
      {
        path: 'alarms',
        loadChildren: () =>
          import('../alarms/alarms.module').then((m) => m.AlarmsModule),
      },
    ],
  },
  {
    path: '',
    pathMatch: 'full',
    redirectTo: 'connections',
  },
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule],
})
export class CoreRoutingModule {}
