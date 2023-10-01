import { NgModule } from '@angular/core';
import { CommonModule, NgOptimizedImage } from '@angular/common';

import { CoreRoutingModule } from './core-routing.module';
import { CoreComponent } from './core/core.component';
import { IconsModule } from '../login/icons/icons.module';
import { HeaderComponent } from './header/header.component';

@NgModule({
  declarations: [CoreComponent, HeaderComponent],
  imports: [CommonModule, CoreRoutingModule, IconsModule, NgOptimizedImage],
})
export class CoreModule {}
