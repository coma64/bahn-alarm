import { Pipe, PipeTransform } from '@angular/core';
import * as dayjs from 'dayjs';
import * as localizedFormat from 'dayjs/plugin/localizedFormat';

dayjs.extend(localizedFormat);

@Pipe({
  name: 'format',
})
export class FormatPipe implements PipeTransform {
  transform(value: boolean | string | dayjs.Dayjs): string {
    if (typeof value === 'boolean') return value ? 'Yes' : 'No';

    if (typeof value === 'string') value = dayjs(value);

    return value.format('LLL');
  }
}
