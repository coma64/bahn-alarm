import { Injectable } from '@angular/core';
import { Observable, of } from 'rxjs';

@Injectable({
  providedIn: 'root',
})
export class NotifyService {
  error(message: string): void {
    alert(message);
  }

  confirm(message: string): Observable<boolean> {
    return of(confirm(message));
  }

  prompt(message: string): Observable<string | undefined> {
    return of(prompt(message) ?? undefined);
  }
}
