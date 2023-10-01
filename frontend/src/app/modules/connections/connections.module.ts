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
import { OverlayModule } from '@angular/cdk/overlay';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { DepartureSelectComponent } from './departure-select/departure-select.component';
import { NgArrayPipesModule } from 'ngx-pipes';
import { StationSearchComponent } from './station-search/station-search.component';
import { RelativeTimeComponent } from './relative-time/relative-time.component';
import { PortalModule } from '@angular/cdk/portal';

@NgModule({
    imports: [
        CommonModule,
        ConnectionsRoutingModule,
        IconsModule,
        SharedModule,
        OverlayModule,
        ReactiveFormsModule,
        FormsModule,
        NgArrayPipesModule,
        PortalModule,
        ConnectionsComponent,
        ConnectionStatsComponent,
        ConnectionListComponent,
        DepartureComponent,
        ToHumanStatusPipe,
        EditConnectionComponent,
        DepartureSelectComponent,
        StationSearchComponent,
        RelativeTimeComponent,
    ],
})
export class ConnectionsModule {}
