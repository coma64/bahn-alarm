import { Component, OnDestroy, OnInit } from '@angular/core';
import { Subject, takeUntil } from 'rxjs';
import { NotifyService } from '../../shared/services/notify.service';
import { PushNotificationSubscriptionService } from '../push-notification-subscription.service';

@Component({
  selector: 'app-core',
  templateUrl: './core.component.html',
  styleUrls: ['./core.component.scss'],
})
export class CoreComponent implements OnInit, OnDestroy {
  private readonly destroy$ = new Subject<void>();
  constructor(
    private readonly pushNotificationSubscription: PushNotificationSubscriptionService,
  ) {}

  ngOnInit(): void {
    this.pushNotificationSubscription
      .register(false)
      .pipe(takeUntil(this.destroy$))
      .subscribe();
  }

  ngOnDestroy(): void {
    this.destroy$.next();
    this.destroy$.complete();
  }
}
