import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { BannerComponent } from './components/banner/banner.component';
import { IconsModule } from '../login/icons/icons.module';
import { NextRelativeTimePipe } from './pipes/next-relative-time.pipe';
import { ToRelativeTimePipe } from './pipes/to-relative-time.pipe';
import { SpinnerComponent } from './components/spinner/spinner.component';
import { IsIncludedInPipe } from './pipes/is-included-in.pipe';
import { FormatPipe } from './pipes/format.pipe';
import { TimeSincePipe } from './pipes/time-since.pipe';

@NgModule({
  declarations: [
    BannerComponent,
    NextRelativeTimePipe,
    ToRelativeTimePipe,
    SpinnerComponent,
    IsIncludedInPipe,
    FormatPipe,
    TimeSincePipe,
  ],
  imports: [CommonModule, IconsModule],
  exports: [
    BannerComponent,
    ToRelativeTimePipe,
    NextRelativeTimePipe,
    SpinnerComponent,
    IsIncludedInPipe,
    FormatPipe,
    TimeSincePipe,
  ],
})
export class SharedModule {}
