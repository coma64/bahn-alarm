import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { ConnectionsRoutingModule } from './connections-routing.module';
import { ConnectionsComponent } from './connections/connections.component';
import { IconsModule } from '../login/icons/icons.module';
import { ConnectionStatsComponent } from './stats/connection-stats.component';
import { SharedModule } from '../shared/shared.module';
import { ConnectionListComponent } from './connection-list/connection-list.component';
import { DepartureComponent } from './departure/departure.component';
import { ToHumanStatusPipe } from './to-human-status.pipe';
import { EditConnectionComponent } from './edit-connection/edit-connection.component';

@NgModule({
  declarations: [ConnectionsComponent, ConnectionStatsComponent, ConnectionListComponent, DepartureComponent, ToHumanStatusPipe, EditConnectionComponent],
  imports: [CommonModule, ConnectionsRoutingModule, IconsModule, SharedModule],
})
export class ConnectionsModule {}
