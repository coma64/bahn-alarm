import { Pipe, PipeTransform } from '@angular/core';
import { RelativeTime } from '../../connections/relative-time/relative-time';

@Pipe({
    name: 'toRelativeTime',
    standalone: true,
})
export class ToRelativeTimePipe implements PipeTransform {
  transform(value: string): RelativeTime {
    return RelativeTime.fromIso(value);
  }
}
