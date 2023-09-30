import {Component, Inject, isDevMode, OnDestroy, OnInit} from '@angular/core';
import { SwUpdate } from '@angular/service-worker';
import {catchError, EMPTY, exhaustMap, filter, Observable, Subject, take, takeUntil, timer} from 'rxjs';
import { fromPromise } from 'rxjs/internal/observable/innerFrom';
import {Select} from "@ngxs/store";
import {UserState, UserStateModel} from "./state/user.state";
import LogRocket from "logrocket";
import Rollbar from "rollbar";
import {rollbarService} from "./rollbar";

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss'],
})
export class AppComponent implements OnInit, OnDestroy {
  @Select(UserState) user!: Observable<UserStateModel>;
  private readonly destroy$ = new Subject<void>();

  constructor(private readonly swUpdate: SwUpdate, @Inject(rollbarService) private readonly rollbar: Rollbar) {
    this.user.pipe(take(1)).subscribe(({user}) => {
      if (user) {
        LogRocket.identify(user.id.toString(), {
          name: user.name,
        });

        this.rollbar.configure({payload: {user: {id: user.id, name: user.name}}})
      }
    })
  }

  ngOnInit(): void {
    // if a new version is available it is handled by the below subscription
    timer(0, 5 * 60_000)
      .pipe(
        exhaustMap(() => fromPromise(this.swUpdate.checkForUpdate())),
        catchError((e: Error) => {
          if (isDevMode()) return EMPTY;
          throw e;
        }),
        takeUntil(this.destroy$),
      )
      .subscribe();

    this.swUpdate.versionUpdates
      .pipe(
        filter((e) => e.type === 'VERSION_READY'),
        takeUntil(this.destroy$),
      )
      .subscribe(
        () =>
          confirm(
            'A new version of this app is available. Do you want to reload?',
          ) && location.reload(),
      );
  }

  ngOnDestroy(): void {
    this.destroy$.next();
    this.destroy$.complete();
  }
}
