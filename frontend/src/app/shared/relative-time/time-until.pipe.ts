import { OnDestroy, Pipe, PipeTransform } from '@angular/core';
import { Subject, takeUntil, timer } from 'rxjs';
import { RelativeTime } from './relative-time';

@Pipe({
  name: 'timeUntil',
  pure: false,
})
export class TimeUntilPipe implements PipeTransform, OnDestroy {
  private readonly next$ = new Subject<void>();
  private currentTime?: RelativeTime;
  private currentTimeDiff = '';
  private alwaysShowMinutes = false;

  ngOnDestroy(): void {
    this.next$.next();
    this.next$.complete();
  }

  transform(time: RelativeTime, alwaysShowMinutes = false): string {
    this.alwaysShowMinutes = alwaysShowMinutes;

    if (time !== this.currentTime) {
      this.currentTime = time;
      this.currentTimeDiff = time?.timeUntil(this.alwaysShowMinutes) ?? '';
      this.setupTimer();
    }

    return this.currentTimeDiff;
  }

  private setupTimer(): void {
    this.next$.next();

    timer(1000)
      .pipe(takeUntil(this.next$))
      .subscribe(
        () =>
          (this.currentTimeDiff =
            this.currentTime?.timeUntil(this.alwaysShowMinutes) ?? ''),
      );
  }
}
