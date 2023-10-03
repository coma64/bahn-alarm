import './init-dayjs';

import LogRocket from 'logrocket';
import { AppComponent } from './app/app.component';
import { ServiceWorkerModule } from '@angular/service-worker';
import { ApiModule, Configuration } from './app/api';
import { NgxsStoragePluginModule } from '@ngxs/storage-plugin';
import {
  NgxsRouterPluginModule,
  RouterDataResolved,
  RouterNavigated,
  RouterNavigation,
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
import { routes } from './app/routes';
import { bootstrapApplication, BrowserModule } from '@angular/platform-browser';
import { JwtExpiredInterceptor } from './app/interceptors/jwt-expired.interceptor';
import {
  HTTP_INTERCEPTORS,
  provideHttpClient,
  withInterceptorsFromDi,
} from '@angular/common/http';
import {
  RollbarErrorHandler,
  rollbarFactory,
  rollbarService,
} from './app/rollbar';
import { ErrorHandler, importProvidersFrom, isDevMode } from '@angular/core';
import { RouterModule } from '@angular/router';

LogRocket.init('qmpbyd/bahn-alarm', {
  console: { isEnabled: { warn: true, error: true } },
});

bootstrapApplication(AppComponent, {
  providers: [
    importProvidersFrom(
      BrowserModule,
      RouterModule.forRoot(routes),
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
