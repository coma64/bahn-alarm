import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { ConnectionsRoutingModule } from './connections-routing.module';
import { ConnectionsComponent } from './connections/connections.component';
import { IconsModule } from '../login/icons/icons.module';
import { ConnectionStatsComponent } from './stats/connection-stats.component';
import { SharedModule } from '../shared/shared.module';

@NgModule({
  declarations: [ConnectionsComponent, ConnectionStatsComponent],
  imports: [CommonModule, ConnectionsRoutingModule, IconsModule, SharedModule],
})
export class ConnectionsModule {}
