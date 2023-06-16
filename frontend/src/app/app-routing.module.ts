import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { isLoggedInGuard } from './guards/logged-in.guard';

const routes: Routes = [
  {
    path: '',
    pathMatch: 'full',
    redirectTo: 'connections',
  },
  {
    path: 'login',
    loadChildren: async () =>
      (await import('./modules/login/login.module')).LoginModule,
  },
  {
    path: '',
    loadChildren: async () =>
      (await import('./modules/core/core.module')).CoreModule,
    canActivate: [isLoggedInGuard],
  },
  {
    path: '**',
    redirectTo: 'connections',
  },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule],
})
export class AppRoutingModule {}
