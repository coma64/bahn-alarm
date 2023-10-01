import { platformBrowserDynamic } from '@angular/platform-browser-dynamic';
import './init-dayjs';

import LogRocket from 'logrocket';
import { AppComponent } from './app/app.component';
import { ServiceWorkerModule } from '@angular/service-worker';
import { ApiModule, Configuration } from './app/api';
import { NgxsStoragePluginModule } from '@ngxs/storage-plugin';
import {
  RouterNavigation,
  RouterDataResolved,
  RouterNavigated,
  NgxsRouterPluginModule,
} from '@ngxs/router-plugin';
import { NgxsLoggerPluginModule } from '@ngxs/logger-plugin';
import { NgxsReduxDevtoolsPluginModule } from '@ngxs/devtools-plugin';
import { environment } from './environments/environment';
import { AlarmsState } from './app/state/alarms.state';
import { AlarmedDevicesState } from './app/state/alarmed-devices.state';
import { ConnectionsState } from './app/state/connections.state';
import { ConnectionStatsState } from './app/state/connection-stats.state';
import { UserState } from './app/state/user.state';
import { NgxsModule } from '@ngxs/store';
import { AppRoutingModule } from './app/app-routing.module';
import { BrowserModule, bootstrapApplication } from '@angular/platform-browser';
import { JwtExpiredInterceptor } from './app/interceptors/jwt-expired.interceptor';
import {
  HTTP_INTERCEPTORS,
  withInterceptorsFromDi,
  provideHttpClient,
} from '@angular/common/http';
import {
  RollbarErrorHandler,
  rollbarService,
  rollbarFactory,
} from './app/rollbar';
import { ErrorHandler, isDevMode, importProvidersFrom } from '@angular/core';

LogRocket.init('qmpbyd/bahn-alarm');

bootstrapApplication(AppComponent, {
  providers: [
    importProvidersFrom(
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
      ApiModule.forRoot(
        () =>
          new Configuration({
            basePath: environment.api,
            withCredentials: true,
          }),
      ),
      ServiceWorkerModule.register('ngsw-worker.js', {
        enabled: !isDevMode(),
        // Register the ServiceWorker as soon as the application is stable
        // or after 30 seconds (whichever comes first).
        registrationStrategy: 'registerWhenStable:30000',
      }),
    ),
    { provide: ErrorHandler, useClass: RollbarErrorHandler },
    { provide: rollbarService, useFactory: rollbarFactory },
    {
      provide: HTTP_INTERCEPTORS,
      multi: true,
      useExisting: JwtExpiredInterceptor,
    },
    provideHttpClient(withInterceptorsFromDi()),
  ],
}).catch((err) => console.error(err));
