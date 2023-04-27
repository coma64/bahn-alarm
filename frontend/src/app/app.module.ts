import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { NgxsModule } from '@ngxs/store';
import { UserState } from './state/user.state';
import { environment } from '../environments/environment';
import { NgxsReduxDevtoolsPluginModule } from '@ngxs/devtools-plugin';
import { NgxsLoggerPluginModule } from '@ngxs/logger-plugin';
import { ApiModule, Configuration } from './api';
import { HTTP_INTERCEPTORS, HttpClientModule } from '@angular/common/http';
import { NgxsRouterPluginModule } from '@ngxs/router-plugin';
import { NgxsStoragePluginModule } from '@ngxs/storage-plugin';
import { JwtExpiredInterceptor } from './interceptors/jwt-expired.interceptor';
import { ConnectionStatsState } from './state/connection-stats.state';
import { ConnectionsState } from './state/connections.state';
import { AlarmedDevicesState } from './state/alarmed-devices.state';

@NgModule({
  declarations: [AppComponent],
  imports: [
    BrowserModule,
    AppRoutingModule,
    NgxsModule.forRoot(
      [UserState, ConnectionStatsState, ConnectionsState, AlarmedDevicesState],
      {
        developmentMode: !environment.production,
      },
    ),
    NgxsReduxDevtoolsPluginModule.forRoot(),
    NgxsLoggerPluginModule.forRoot(),
    NgxsRouterPluginModule.forRoot(),
    NgxsStoragePluginModule.forRoot({
      key: [UserState],
    }),
    HttpClientModule,
    ApiModule.forRoot(
      () =>
        new Configuration({ basePath: environment.api, withCredentials: true }),
    ),
  ],
  providers: [
    {
      provide: HTTP_INTERCEPTORS,
      multi: true,
      useExisting: JwtExpiredInterceptor,
    },
  ],
  bootstrap: [AppComponent],
})
export class AppModule {}
