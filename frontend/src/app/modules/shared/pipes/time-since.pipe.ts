import { Pipe, PipeTransform } from '@angular/core';
import dayjs from 'dayjs/esm';
import { interval, map, Observable, startWith } from 'rxjs';

@Pipe({
  name: 'timeSince',
})
export class TimeSincePipe implements PipeTransform {
  transform(value: dayjs.Dayjs | string): Observable<string> {
    const date = typeof value === 'string' ? dayjs(value) : value;

    return interval(1000).pipe(
      startWith(date.from(dayjs())),
      map(() => date.from(dayjs())),
    );
  }
}
