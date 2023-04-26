import { NgModule } from '@angular/core';
import { CommonModule, NgOptimizedImage } from '@angular/common';

import { CoreRoutingModule } from './core-routing.module';
import { CoreComponent } from './core/core.component';
import { NavComponent } from './nav/nav.component';
import { IconsModule } from '../login/icons/icons.module';

@NgModule({
  declarations: [CoreComponent, NavComponent],
  imports: [CommonModule, CoreRoutingModule, IconsModule, NgOptimizedImage],
})
export class CoreModule {}
