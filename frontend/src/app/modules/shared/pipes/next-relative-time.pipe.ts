import { Pipe, PipeTransform } from '@angular/core';
import { interval, map, Observable, startWith } from 'rxjs';
import {
  RelativeTime,
  TimeUntilOptions,
} from '../../connections/relative-time/relative-time';

@Pipe({
  name: 'nextRelativeTime',
})
export class NextRelativeTimePipe implements PipeTransform {
  transform(
    time: RelativeTime,
    options: TimeUntilOptions = RelativeTime.TIME_UNTIL_DEFAULTS,
  ): Observable<string> {
    return interval(1000).pipe(
      startWith(time.timeUntil(options)),
      map(() => time.timeUntil(options)),
    );
  }
}
