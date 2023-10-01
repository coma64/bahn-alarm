import { NgModule } from '@angular/core';
import { CommonModule, NgOptimizedImage } from '@angular/common';

import { CoreRoutingModule } from './core-routing.module';
import { CoreComponent } from './core/core.component';
import { IconsModule } from '../login/icons/icons.module';
import { HeaderComponent } from './header/header.component';

@NgModule({
    imports: [CommonModule, CoreRoutingModule, IconsModule, NgOptimizedImage, CoreComponent, HeaderComponent],
})
export class CoreModule {}
