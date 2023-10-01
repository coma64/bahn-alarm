import { Pipe, PipeTransform } from '@angular/core';
import { interval, map, Observable, startWith } from 'rxjs';
import {
  RelativeTime,
  TimeUntilOptions,
} from '../../connections/relative-time/relative-time';

@Pipe({
    name: 'nextRelativeTime',
    standalone: true,
})
export class NextRelativeTimePipe implements PipeTransform {
  transform(
    time: RelativeTime,
    options: TimeUntilOptions = RelativeTime.timeUntilDefaults,
  ): Observable<string> {
    return interval(1000).pipe(
      startWith(time.timeUntil(options)),
      map(() => time.timeUntil(options)),
    );
  }
}
