import { Injectable } from '@angular/core';
import {
  HttpRequest,
  HttpHandler,
  HttpEvent,
  HttpInterceptor,
  HttpErrorResponse,
} from '@angular/common/http';
import { catchError, EMPTY, Observable } from 'rxjs';
import { Store } from '@ngxs/store';
import { Navigate } from '@ngxs/router-plugin';
import { User } from '../state/user.actions';

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

    this.store.dispatch(new User.Logout());
    return EMPTY;
  }
}
