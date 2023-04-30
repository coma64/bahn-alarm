import { NgModule } from '@angular/core';
import {
  AlertTriangle,
  Check,
  LogIn,
  Plus,
  Trash,
  ArrowLeft,
  Bell,
  Eye,
  User,
  LogOut,
  X,
  RefreshCw,
  ChevronDown,
} from 'angular-feather/icons';
import { FeatherModule } from 'angular-feather';

const icons = {
  LogIn,
  AlertTriangle,
  Plus,
  Trash,
  Check,
  ArrowLeft,
  Bell,
  Eye,
  User,
  LogOut,
  X,
  RefreshCw,
  ChevronDown,
};

@NgModule({
  imports: [FeatherModule.pick(icons)],
  exports: [FeatherModule],
})
export class IconsModule {}
