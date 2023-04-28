import { OnDestroy, Pipe, PipeTransform } from '@angular/core';
import { Subject, takeUntil, timer } from 'rxjs';
import {
  RelativeTime,
  TimeUntilOptions,
} from '../../connections/relative-time/relative-time';

@Pipe({
  name: 'timeUntil',
  pure: false,
})
export class TimeUntilPipe implements PipeTransform, OnDestroy {
  private readonly next$ = new Subject<void>();
  private currentTime?: RelativeTime;
  private currentTimeDiff = '';
  private options: TimeUntilOptions = RelativeTime.TIME_UNTIL_DEFAULTS;

  ngOnDestroy(): void {
    this.next$.next();
    this.next$.complete();
  }

  transform(
    time: RelativeTime,
    options: TimeUntilOptions = RelativeTime.TIME_UNTIL_DEFAULTS,
  ): string {
    this.options = options;

    if (time !== this.currentTime) {
      this.currentTime = time;
      this.currentTimeDiff = time?.timeUntil(options) ?? '';
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
            this.currentTime?.timeUntil(this.options) ?? ''),
      );
  }
}
