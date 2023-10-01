import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { LoginRoutingModule } from './login-routing.module';
import { LoginComponent } from './login/login.component';
import { IconsModule } from './icons/icons.module';
import { ReactiveFormsModule } from '@angular/forms';
import { SharedModule } from '../shared/shared.module';

@NgModule({
    imports: [
        CommonModule,
        LoginRoutingModule,
        IconsModule,
        ReactiveFormsModule,
        SharedModule,
        LoginComponent,
    ],
})
export class LoginModule {}
