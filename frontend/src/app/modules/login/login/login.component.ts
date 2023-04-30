import { Component, OnDestroy } from '@angular/core';
import { FormBuilder, Validators } from '@angular/forms';
import { Subject, takeUntil } from 'rxjs';
import { Store } from '@ngxs/store';
import { UserActions } from '../../../state/user.actions';
import { AuthService, LoginRequest } from '../../../api';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss'],
})
export class LoginComponent implements OnDestroy {
  readonly form = this.fb.nonNullable.group({
    username: ['', Validators.required],
    password: ['', Validators.required],
  });

  isInvalid = false;

  private readonly destroy$ = new Subject<void>();

  constructor(
    private readonly fb: FormBuilder,
    private readonly store: Store,
    private readonly auth: AuthService,
  ) {}

  ngOnDestroy(): void {
    this.destroy$.next();
    this.destroy$.complete();
  }

  onSubmit(): void {
    if (this.form.invalid) return;

    this.auth
      .authLoginPost(this.form.value as LoginRequest)
      .pipe(takeUntil(this.destroy$))
      .subscribe({
        next: (user) => this.store.dispatch(new UserActions.LoginSuccess(user)),
        error: () => {
          this.isInvalid = true;
        },
      });
  }
}
