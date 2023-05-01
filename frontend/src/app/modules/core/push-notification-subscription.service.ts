import { Injectable } from '@angular/core';
import { SwPush } from '@angular/service-worker';
import {
  NotificationsService,
  PushNotificationSubscription,
  RawSubscription,
} from '../../api';
import { NotifyService } from '../shared/services/notify.service';
import { EMPTY, forkJoin, Observable, switchMap, tap } from 'rxjs';
import { fromPromise } from 'rxjs/internal/observable/innerFrom';
import { Store } from '@ngxs/store';
import { UserState, UserStateModel } from '../../state/user.state';
import { UserActions } from '../../state/user.actions';

@Injectable({
  providedIn: 'root',
})
export class PushNotificationSubscriptionService {
  get isRegistered(): boolean {
    return this.swPush.isEnabled;
  }

  constructor(
    private readonly swPush: SwPush,
    private readonly notifications: NotificationsService,
    private readonly notify: NotifyService,
    private readonly store: Store,
  ) {
    this.swPush.messages.subscribe((message) => console.log({ message }));
    this.swPush.notificationClicks.subscribe((click) => console.log({ click }));
  }

  register(
    ignorePreviousDenials: boolean,
  ): Observable<PushNotificationSubscription> {
    if (
      !ignorePreviousDenials &&
      this.store.selectSnapshot<UserStateModel>(UserState)
        .hasDeniedPushNotifications
    )
      return EMPTY;

    return this.notify
      .confirm(
        'Do you want to setup push notification for delayed connections?',
      )
      .pipe(
        switchMap((isConfirmed) => {
          if (isConfirmed)
            return this.notifications.notificationsVapidKeysGet();

          this.store.dispatch(UserActions.DeniedPushNotification);
          return EMPTY;
        }),
        switchMap(({ publicKey }) =>
          forkJoin([
            this.notify.prompt('How do you want to name this device?'),
            this.requestSubscription(publicKey),
          ]),
        ),
        switchMap(([name, subscription]) =>
          this.notifications.notificationsPushSubscriptionsPost({
            name: name ?? '',
            subscription: subscription.toJSON() as RawSubscription,
          }),
        ),
        tap({
          error: (err: Error) =>
            this.notify.error(
              `An unknown error occurred while setting up the push notification subscription.\n\n${
                err.stack ?? err.message
              }`,
            ),
        }),
      );
  }

  private requestSubscription(publicKey: string): Observable<PushSubscription> {
    return fromPromise(
      this.swPush.requestSubscription({ serverPublicKey: publicKey }),
    );
  }
}
