import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { AlarmsComponent } from './alarms/alarms.component';

const routes: Routes = [
  {
    path: '',
    component: AlarmsComponent,
  },
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule],
})
export class AlarmsRoutingModule {}
