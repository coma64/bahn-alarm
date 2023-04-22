import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { BannerComponent } from './banner/banner.component';
import { IconsModule } from '../login/icons/icons.module';
import { TimeUntilPipe } from './relative-time/time-until.pipe';
import { ToRelativeTimePipe } from './relative-time/to-relative-time.pipe';

@NgModule({
  declarations: [BannerComponent, TimeUntilPipe, ToRelativeTimePipe],
  imports: [CommonModule, IconsModule],
  exports: [BannerComponent, ToRelativeTimePipe, TimeUntilPipe],
})
export class SharedModule {}
