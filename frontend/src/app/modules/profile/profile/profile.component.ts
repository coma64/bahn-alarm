import { Component, OnInit } from '@angular/core';
import { Select, Store } from '@ngxs/store';
import { State } from '../../../state/state';
import { Observable } from 'rxjs';
import { AuthService, NotificationsService } from '../../../api';
import { AlarmedDeviceActions } from '../../../state/alarmed-devices.actions';
import { UserActions } from '../../../state/user.actions';
import { NotifyService } from '../../shared/services/notify.service';

@Component({
  selector: 'app-profile',
  templateUrl: './profile.component.html',
  styleUrls: ['./profile.component.scss'],
})
export class ProfileComponent implements OnInit {
  @Select() user$!: Observable<State['user']>;

  constructor(
    private readonly notifications: NotificationsService,
    private readonly auth: AuthService,
    private readonly store: Store,
    private readonly notify: NotifyService,
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

  onLogout(): void {
    this.auth.authLogoutPost().subscribe({
      next: () => this.store.dispatch(new UserActions.Logout()),
      error: () =>
        this.notify.error(
          'An unknown error occurred while logging you out. Please try again',
        ),
    });
  }
}
