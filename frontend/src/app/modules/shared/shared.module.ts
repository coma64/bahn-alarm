import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { BannerComponent } from './banner/banner.component';
import { IconsModule } from '../login/icons/icons.module';
import { TimeUntilPipe } from './relative-time/time-until.pipe';
import { ToRelativeTimePipe } from './relative-time/to-relative-time.pipe';
import { SpinnerComponent } from './spinner/spinner.component';
import { StationSearchComponent } from './station-search/station-search.component';
import { ReactiveFormsModule } from '@angular/forms';
import { PortalModule } from '@angular/cdk/portal';
import { RelativeTimeComponent } from './relative-time/relative-time.component';
import { IsIncludedInPipe } from './is-included-in.pipe';

@NgModule({
  declarations: [
    BannerComponent,
    TimeUntilPipe,
    ToRelativeTimePipe,
    SpinnerComponent,
    StationSearchComponent,
    RelativeTimeComponent,
    IsIncludedInPipe,
  ],
  imports: [CommonModule, IconsModule, ReactiveFormsModule, PortalModule],
  exports: [
    BannerComponent,
    ToRelativeTimePipe,
    TimeUntilPipe,
    SpinnerComponent,
    StationSearchComponent,
    RelativeTimeComponent,
    IsIncludedInPipe,
  ],
})
export class SharedModule {}
