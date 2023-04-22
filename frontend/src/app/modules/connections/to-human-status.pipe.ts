import { Pipe, PipeTransform } from '@angular/core';
import { TrackedDeparture } from '../../api';

@Pipe({
  name: 'toHumanStatus',
})
export class ToHumanStatusPipe implements PipeTransform {
  transform(departure: TrackedDeparture): string {
    const status = departure.status;
    if (status === 'on-time') return 'on time';
    else if (status === 'delayed') return `+${departure.delay}m delay`;
    else if (status === 'canceled') return 'canceled';
    else return 'not checked';
  }
}
