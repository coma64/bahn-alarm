import { Component, OnDestroy } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, Validators } from '@angular/forms';
import { Subject, takeUntil } from 'rxjs';
import { Store } from '@ngxs/store';
import { LoginSuccess } from '../../../../state/user.actions';

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
    private readonly http: HttpClient,
    private readonly fb: FormBuilder,
    private readonly store: Store
  ) {}

  ngOnDestroy(): void {
    this.destroy$.next();
    this.destroy$.complete();
  }

  onSubmit(): void {
    if (this.form.invalid) return;

    this.http
      .post('http://localhost:8090/auth/login', this.form.value)
      .pipe(takeUntil(this.destroy$))
      .subscribe({
        next: () =>
          this.store.dispatch(
            new LoginSuccess(this.form.controls.username.value)
          ),
        error: () => {
          this.isInvalid = true;
        },
      });
  }
}
