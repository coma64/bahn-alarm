import {NgModule, isDevMode, ErrorHandler} from '@angular/core';
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
import {
  NgxsRouterPluginModule,
  RouterDataResolved,
  RouterNavigated,
  RouterNavigation,
} from '@ngxs/router-plugin';
import { NgxsStoragePluginModule } from '@ngxs/storage-plugin';
import { JwtExpiredInterceptor } from './interceptors/jwt-expired.interceptor';
import { ConnectionStatsState } from './state/connection-stats.state';
import { ConnectionsState } from './state/connections.state';
import { AlarmedDevicesState } from './state/alarmed-devices.state';
import { AlarmsState } from './state/alarms.state';
import { ServiceWorkerModule } from '@angular/service-worker';
import {RollbarErrorHandler, rollbarFactory, rollbarService} from "./rollbar";

@NgModule({
  declarations: [AppComponent],
  imports: [
    BrowserModule,
    AppRoutingModule,
    NgxsModule.forRoot(
      [
        UserState,
        ConnectionStatsState,
        ConnectionsState,
        AlarmedDevicesState,
        AlarmsState,
      ],
      {
        developmentMode: !environment.production,
      },
    ),
    NgxsReduxDevtoolsPluginModule.forRoot(),
    NgxsLoggerPluginModule.forRoot({
      collapsed: true,
      filter: (action) => {
        return !(
          action instanceof RouterNavigation ||
          action instanceof RouterDataResolved ||
          action instanceof RouterNavigated
        );
      },
      disabled: environment.production,
    }),
    NgxsRouterPluginModule.forRoot(),
    NgxsStoragePluginModule.forRoot({
      key: [UserState],
    }),
    HttpClientModule,
    ApiModule.forRoot(
      () =>
        new Configuration({ basePath: environment.api, withCredentials: true }),
    ),
    ServiceWorkerModule.register('ngsw-worker.js', {
      enabled: !isDevMode(),
      // Register the ServiceWorker as soon as the application is stable
      // or after 30 seconds (whichever comes first).
      registrationStrategy: 'registerWhenStable:30000',
    }),
  ],
  providers: [
    { provide: ErrorHandler, useClass: RollbarErrorHandler },
    { provide: rollbarService, useFactory: rollbarFactory },
    {
      provide: HTTP_INTERCEPTORS,
      multi: true,
      useExisting: JwtExpiredInterceptor,
    },
  ],
  bootstrap: [AppComponent],
})
export class AppModule {}
