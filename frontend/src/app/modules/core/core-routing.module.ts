import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

const routes: Routes = [
  {
    path: 'connections',
    loadChildren: () =>
      import('../connections/connections.module').then(
        (m) => m.ConnectionsModule,
      ),
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
