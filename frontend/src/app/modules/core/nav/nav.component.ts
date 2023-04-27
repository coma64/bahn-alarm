import { Component } from '@angular/core';
import { AuthService } from '../../../api';
import { Store } from '@ngxs/store';
import { UserActions } from '../../../state/user.actions';

@Component({
  selector: 'app-nav',
  templateUrl: './nav.component.html',
  styleUrls: ['./nav.component.scss'],
})
export class NavComponent {
  constructor(
    private readonly auth: AuthService,
    private readonly store: Store,
  ) {}

  onLogout(): void {
    this.auth.authLogoutPost().subscribe({
      next: () => this.store.dispatch(new UserActions.Logout()),
      error: () =>
        alert(
          'An unknown error occurred while logging you out. Please try again',
        ),
    });
  }
}
