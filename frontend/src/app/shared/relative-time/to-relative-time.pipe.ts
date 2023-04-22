import { Pipe, PipeTransform } from '@angular/core';
import { RelativeTime } from './relative-time';

@Pipe({
  name: 'toRelativeTime',
})
export class ToRelativeTimePipe implements PipeTransform {
  transform(value: string): RelativeTime {
    return RelativeTime.fromIso(value);
  }
}
