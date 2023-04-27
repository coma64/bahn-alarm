import { Injectable } from '@angular/core';
import {
  HttpErrorResponse,
  HttpEvent,
  HttpHandler,
  HttpInterceptor,
  HttpRequest,
} from '@angular/common/http';
import { catchError, EMPTY, Observable } from 'rxjs';
import { Store } from '@ngxs/store';
import { UserActions } from '../state/user.actions';

@Injectable({
  providedIn: 'root',
})
export class JwtExpiredInterceptor implements HttpInterceptor {
  constructor(private readonly store: Store) {}

  intercept(
    request: HttpRequest<unknown>,
    next: HttpHandler,
  ): Observable<HttpEvent<unknown>> {
    return next
      .handle(request)
      .pipe(catchError((err) => this.handleHttpError(err)));
  }

  private handleHttpError(err: HttpErrorResponse): Observable<any> {
    if (err.status !== 401) throw err;

    this.store.dispatch(new UserActions.Logout());
    return EMPTY;
  }
}
