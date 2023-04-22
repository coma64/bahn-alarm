import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { BannerComponent } from './banner/banner.component';
import { IconsModule } from '../login/icons/icons.module';
import { TimeUntilPipe } from './relative-time/time-until.pipe';
import { ToRelativeTimePipe } from './relative-time/to-relative-time.pipe';
import { SpinnerComponent } from './spinner/spinner.component';

@NgModule({
  declarations: [
    BannerComponent,
    TimeUntilPipe,
    ToRelativeTimePipe,
    SpinnerComponent,
  ],
  imports: [CommonModule, IconsModule],
  exports: [
    BannerComponent,
    ToRelativeTimePipe,
    TimeUntilPipe,
    SpinnerComponent,
  ],
})
export class SharedModule {}
