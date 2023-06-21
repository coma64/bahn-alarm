import { Injectable, OnDestroy } from '@angular/core';
import { SwPush } from '@angular/service-worker';
import {
  NotificationsService,
  PushNotificationSubscription,
  RawSubscription,
} from '../../api';
import { NotifyService } from '../shared/services/notify.service';
import {
  EMPTY,
  forkJoin,
  Observable,
  of,
  Subject,
  switchMap,
  take,
  takeUntil,
  tap,
} from 'rxjs';
import { fromPromise } from 'rxjs/internal/observable/innerFrom';
import { Store } from '@ngxs/store';
import { UserActions } from '../../state/user.actions';

@Injectable({
  providedIn: 'root',
})
export class PushNotificationSubscriptionService implements OnDestroy {
  private readonly destroy$ = new Subject<void>();
  private _isRegistered = true;

  get isRegistered(): boolean {
    return this._isRegistered;
  }

  constructor(
    private readonly swPush: SwPush,
    private readonly notifications: NotificationsService,
    private readonly notify: NotifyService,
    private readonly store: Store,
  ) {
    this.swPush.subscription
      .pipe(takeUntil(this.destroy$))
      .subscribe((sub) => (this._isRegistered = !!sub));
  }

  ngOnDestroy(): void {
    this.destroy$.next();
    this.destroy$.complete();
  }

  askUserAndRegister(): Observable<PushNotificationSubscription> {
    return this.notify
      .confirm(
        'Do you want to setup push notifications, to receive notification about delayed connections?',
      )
      .pipe(
        switchMap((isConfirmed) => (isConfirmed ? this.register() : EMPTY)),
      );
  }

  register(): Observable<PushNotificationSubscription> {
    return this.notify.prompt('How do you want to name this device?').pipe(
      switchMap((name) => forkJoin([of(name), this.requestSubscription()])),
      switchMap(([name, subscription]) =>
        this.notifications.notificationsPushSubscriptionsPost({
          name: name ?? '',
          subscription: subscription.toJSON() as RawSubscription,
        }),
      ),
      tap({
        next: ({ id }) =>
          this.store.dispatch(new UserActions.RegisteredPushNotifications(id)),
        error: (err: Error) =>
          this.notify.error(
            `An unknown error occurred while setting up the push notification subscription.\n\n${
              err.stack ?? err.message
            }`,
          ),
      }),
      take(1),
    );
  }

  unregister(pushSubId: number): Observable<unknown> {
    if (!this.isRegistered) return EMPTY;

    return forkJoin([
      fromPromise(this.swPush.unsubscribe()),
      this.notifications.notificationsPushSubscriptionsIdDelete(pushSubId),
    ]);
  }

  private requestSubscription(): Observable<PushSubscription> {
    return this.notifications
      .notificationsVapidKeysGet()
      .pipe(
        switchMap(({ publicKey }) =>
          fromPromise(
            this.swPush.requestSubscription({ serverPublicKey: publicKey }),
          ),
        ),
      );
  }
}
