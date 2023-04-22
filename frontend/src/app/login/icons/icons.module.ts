import { NgModule } from '@angular/core';
import { LogIn, AlertTriangle } from 'angular-feather/icons';
import { FeatherModule } from 'angular-feather';

const icons = {
  LogIn,
  AlertTriangle,
};

@NgModule({
  imports: [FeatherModule.pick(icons)],
  exports: [FeatherModule],
})
export class IconsModule {}
