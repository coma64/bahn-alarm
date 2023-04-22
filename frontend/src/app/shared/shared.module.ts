import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { BannerComponent } from './banner/banner.component';
import { IconsModule } from '../login/icons/icons.module';

@NgModule({
  declarations: [BannerComponent],
  imports: [CommonModule, IconsModule],
  exports: [BannerComponent],
})
export class SharedModule {}
