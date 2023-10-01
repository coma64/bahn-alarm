import { Pipe, PipeTransform } from '@angular/core';
import dayjs from 'dayjs/esm';

@Pipe({
    name: 'format',
    standalone: true,
})
export class FormatPipe implements PipeTransform {
  transform(value: boolean | string | dayjs.Dayjs): string {
    if (typeof value === 'boolean') return value ? 'Yes' : 'No';

    if (typeof value === 'string') value = dayjs(value);

    return value.format('LLL');
  }
}
