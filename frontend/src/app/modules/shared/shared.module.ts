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
import { DropdownComponent } from './dropdown/dropdown.component';
import { PortalModule } from '@angular/cdk/portal';
import { NgObjectPipesModule } from 'ngx-pipes';
import { PaginationComponent } from './components/pagination/pagination.component';
import { ReactiveFormsModule } from '@angular/forms';

@NgModule({
  declarations: [
    BannerComponent,
    NextRelativeTimePipe,
    ToRelativeTimePipe,
    SpinnerComponent,
    IsIncludedInPipe,
    FormatPipe,
    TimeSincePipe,
    DropdownComponent,
    PaginationComponent,
  ],
  imports: [
    CommonModule,
    IconsModule,
    PortalModule,
    NgObjectPipesModule,
    ReactiveFormsModule,
  ],
  exports: [
    BannerComponent,
    ToRelativeTimePipe,
    NextRelativeTimePipe,
    SpinnerComponent,
    IsIncludedInPipe,
    FormatPipe,
    TimeSincePipe,
    DropdownComponent,
    PaginationComponent,
  ],
})
export class SharedModule {}
