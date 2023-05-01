import { Component, OnDestroy, OnInit } from '@angular/core';
import { Select, Store } from '@ngxs/store';
import { State } from '../../../state/state';
import { Observable, Subject, takeUntil } from 'rxjs';
import { AuthService, NotificationsService } from '../../../api';
import { AlarmedDeviceActions } from '../../../state/alarmed-devices.actions';
import { UserActions } from '../../../state/user.actions';
import { NotifyService } from '../../shared/services/notify.service';
import { PushNotificationSubscriptionService } from '../../core/push-notification-subscription.service';

@Component({
  selector: 'app-profile',
  templateUrl: './profile.component.html',
  styleUrls: ['./profile.component.scss'],
})
export class ProfileComponent implements OnInit, OnDestroy {
  @Select() user$!: Observable<State['user']>;
  private readonly destroy$ = new Subject<void>();

  constructor(
    private readonly notifications: NotificationsService,
    private readonly auth: AuthService,
    private readonly store: Store,
    private readonly notify: NotifyService,
    protected readonly pushNotificationSubscription: PushNotificationSubscriptionService,
  ) {}

  ngOnInit(): void {
    this.notifications
      .notificationsPushSubscriptionsGet()
      .subscribe((notifications) =>
        this.store.dispatch(
          new AlarmedDeviceActions.Fetched(notifications.subscriptions),
        ),
      );
  }

  ngOnDestroy(): void {
    this.destroy$.next();
    this.destroy$.complete();
  }

  onLogout(): void {
    this.auth.authLogoutPost().subscribe({
      next: () => this.store.dispatch(new UserActions.Logout()),
      error: () =>
        this.notify.error(
          'An unknown error occurred while logging you out. Please try again',
        ),
    });
  }

  onRegisterPushNotifications(): void {
    this.pushNotificationSubscription
      .register()
      .pipe(takeUntil(this.destroy$))
      .subscribe();
  }
}
