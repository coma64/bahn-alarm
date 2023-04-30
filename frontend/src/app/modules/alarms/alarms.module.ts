import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { AlarmsRoutingModule } from './alarms-routing.module';
import { AlarmsComponent } from './alarms/alarms.component';
import { AlarmsListComponent } from './alarms-list/alarms-list.component';
import { SharedModule } from '../shared/shared.module';
import { IconsModule } from '../login/icons/icons.module';
import { ConnectionInfoComponent } from './connection-info/connection-info.component';
import { FiltersComponent } from './filters/filters.component';

@NgModule({
  declarations: [AlarmsComponent, AlarmsListComponent, ConnectionInfoComponent, FiltersComponent],
  imports: [CommonModule, AlarmsRoutingModule, SharedModule, IconsModule],
})
export class AlarmsModule {}
