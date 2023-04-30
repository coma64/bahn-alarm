import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { ProfileRoutingModule } from './profile-routing.module';
import { ProfileComponent } from './profile/profile.component';
import { SharedModule } from '../shared/shared.module';
import { IconsModule } from '../login/icons/icons.module';
import { AlarmedDevicesListComponent } from './alarmed-devices-list/alarmed-devices-list.component';

@NgModule({
  declarations: [ProfileComponent, AlarmedDevicesListComponent],
  imports: [CommonModule, ProfileRoutingModule, SharedModule, IconsModule],
})
export class ProfileModule {}