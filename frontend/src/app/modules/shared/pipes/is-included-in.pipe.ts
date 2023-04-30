import { Pipe, PipeTransform } from '@angular/core';

@Pipe({
  name: 'isIncludedIn',
})
export class IsIncludedInPipe implements PipeTransform {
  transform<T>(value: T, arr: readonly T[]): boolean {
    return arr.includes(value);
  }
}
