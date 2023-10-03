import { inject, Injectable, OnDestroy } from '@angular/core';
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
import { combineLatest } from 'rxjs';

@Injectable({
  providedIn: 'root',
})
export class PushNotificationSubscriptionService implements OnDestroy {
  private readonly destroy$ = new Subject<void>();
  private _isRegistered = false;
  private vapidPublicKey?: string;
  private readonly swPush?: SwPush;

  get isRegistered(): boolean {
    return this._isRegistered;
  }

  constructor(
    private readonly notifications: NotificationsService,
    private readonly notify: NotifyService,
    private readonly store: Store,
  ) {
    try {
      this.swPush = inject(SwPush);
    } catch (e) {
      console.error("Failed to initialize 'SwPush'", e);
    }

    this.swPush?.subscription
      .pipe(takeUntil(this.destroy$))
      .subscribe((sub) => (this._isRegistered = !!sub));

    this.notifications
      .notificationsVapidKeysGet()
      .pipe(takeUntil(this.destroy$))
      .subscribe(({ publicKey }) => (this.vapidPublicKey = publicKey));
  }

  ngOnDestroy(): void {
    this.destroy$.next();
    this.destroy$.complete();
  }

  register(): Observable<PushNotificationSubscription> {
    return this.requestSubscription()
      .pipe(
        switchMap((subscription) =>
          forkJoin([
            this.notify.prompt('How do you want to name this device?'),
            of(subscription),
          ]),
        ),
      )
      .pipe(
        switchMap(([name, subscription]) =>
          this.notifications.notificationsPushSubscriptionsPost({
            name: name ?? '',
            subscription: subscription.toJSON() as RawSubscription,
          }),
        ),
        tap({
          next: ({ id }) =>
            this.store.dispatch(
              new UserActions.RegisteredPushNotifications(id),
            ),
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

  unregister(pushSubId?: number): Observable<unknown> {
    return combineLatest([
      this.isRegistered && this.swPush
        ? fromPromise(this.swPush.unsubscribe())
        : EMPTY,
      !pushSubId
        ? EMPTY
        : this.notifications.notificationsPushSubscriptionsIdDelete(pushSubId),
    ]);
  }

  private requestSubscription(): Observable<PushSubscription> {
    if (!this.vapidPublicKey) {
      this.notify.error("Vapid public key wasn't loaded in advance. Exiting");
      return EMPTY;
    }

    if (!this.swPush) {
      return EMPTY;
    }

    return fromPromise(
      this.swPush.requestSubscription({ serverPublicKey: this.vapidPublicKey }),
    );
  }
}
