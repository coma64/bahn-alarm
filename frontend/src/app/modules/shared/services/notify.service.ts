import { Injectable } from '@angular/core';
import { Observable, of } from 'rxjs';

@Injectable({
  providedIn: 'root',
})
export class NotifyService {
  error(message: string): void {
    // eslint-disable-next-line no-use-before-define
    alert(message);
  }

  confirm(message: string): Observable<boolean> {
    // eslint-disable-next-line no-use-before-define
    return of(confirm(message));
  }
}
