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
        loadChildren: async () =>
          (await import('../connections/connections.module')).ConnectionsModule,
      },
      {
        path: 'profile',
        loadChildren: async () =>
          (await import('../profile/profile.module')).ProfileModule,
      },
      {
        path: 'alarms',
        loadChildren: async () =>
          (await import('../alarms/alarms.module')).AlarmsModule,
      },
    ],
  },
  {
    path: '',
    pathMatch: 'full',
    redirectTo: '/connections',
  },
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule],
})
export class CoreRoutingModule {}
