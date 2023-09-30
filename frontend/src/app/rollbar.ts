import Rollbar, {Configuration} from 'rollbar';

import {
  Injectable,
  Inject,
  InjectionToken,
  ErrorHandler
} from '@angular/core';
import LogRocket from "logrocket";

const rollbarConfig: Configuration = {
  accessToken: '771442c822394e9787f6fc2dfcaa645a',
  captureUncaught: true,
  captureUnhandledRejections: true,
};

export const rollbarService = new InjectionToken<Rollbar>('rollbar');

@Injectable()
export class RollbarErrorHandler implements ErrorHandler {
  constructor(@Inject(rollbarService) private readonly rollbar: Rollbar) {}

  handleError(err:any) : void {
    this.rollbar.error(err.originalError || err);
    LogRocket.captureException(err);
    console.error(err);
  }
}

export function rollbarFactory() {
  return new Rollbar(rollbarConfig);
}
