import { Pipe, PipeTransform } from '@angular/core';
import { TrackedDeparture } from '../../api';

@Pipe({
  name: 'toHumanStatus',
})
export class ToHumanStatusPipe implements PipeTransform {
  transform(departure: TrackedDeparture): string {
    const {status} = departure;
    if (status === 'on-time') return 'on time';
    if (status === 'delayed') return `+${departure.delay}m delay`;
    if (status === 'canceled') return 'canceled';
    return 'not checked';
  }
}
