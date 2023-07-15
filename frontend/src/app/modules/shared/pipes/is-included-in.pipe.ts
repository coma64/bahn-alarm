import { Pipe, PipeTransform } from '@angular/core';

@Pipe({
  name: 'isIncludedIn',
})
export class IsIncludedInPipe implements PipeTransform {
  transform<T>(value: T, arr: readonly T[], comparingKey?: keyof T): boolean {
    if (comparingKey !== undefined)
      return !!arr.find((e) => e[comparingKey] === value[comparingKey]);
    return arr.includes(value);
  }
}
