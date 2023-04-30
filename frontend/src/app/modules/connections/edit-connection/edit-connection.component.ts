import { Component, OnDestroy, OnInit } from '@angular/core';
import { RelativeTime } from '../relative-time/relative-time';
import { FormBuilder, Validators } from '@angular/forms';
import { BahnPlace, TrackingService } from '../../../api';
import { Store } from '@ngxs/store';
import { Navigate } from '@ngxs/router-plugin';
import { HttpErrorResponse } from '@angular/common/http';
import { Subject, takeUntil } from 'rxjs';
import { NotifyService } from '../../shared/services/notify.service';
import { Connections } from '../../../state/connections.actions';

@Component({
  selector: 'app-edit-connection',
  templateUrl: './edit-connection.component.html',
  styleUrls: ['./edit-connection.component.scss'],
})
export class EditConnectionComponent implements OnInit, OnDestroy {
  readonly stationForm = this.fb.nonNullable.group({
    from: [undefined as BahnPlace | undefined, Validators.required],
    to: [undefined as BahnPlace | undefined, Validators.required],
  });

  selectedDepartures: readonly RelativeTime[] = [];
  hasTriedSubmitting = false;

  get from(): BahnPlace | undefined {
    return this.stationForm.value.from;
  }

  get to(): BahnPlace | undefined {
    return this.stationForm.value.to;
  }

  private readonly destroy$ = new Subject<void>();

  constructor(
    private readonly fb: FormBuilder,
    private readonly tracking: TrackingService,
    private readonly notify: NotifyService,
    protected readonly store: Store,
  ) {}

  ngOnInit(): void {
    this.stationForm.valueChanges
      .pipe(takeUntil(this.destroy$))
      .subscribe(() => (this.selectedDepartures = []));
  }

  ngOnDestroy() {
    this.destroy$.next();
    this.destroy$.complete();
  }

  onSubmit(): void {
    this.hasTriedSubmitting = true;

    if (!this.from || !this.to) return;

    if (
      !this.selectedDepartures.length &&
      this.notify.confirm(
        "You won't get alarms for this connection because you didn't select any departures. Keep editing?",
      )
    )
      return;

    // TODO: store pagination issues
    this.tracking
      .trackingConnectionsPost({
        from: {
          id: this.from.stationId,
          name: this.from.name,
        },
        to: {
          id: this.to.stationId,
          name: this.to.name,
        },
        departures: this.selectedDepartures.map((d) => ({
          departure: d.toIsoZeroBased(),
        })),
      })
      .subscribe({
        next: (connection) =>
          this.store.dispatch([
            new Connections.Created(connection),
            new Navigate(['/connections']),
          ]),
        error: (err: HttpErrorResponse) => {
          if (err.status === 409)
            this.notify.error(
              "You're already tracking this connection. You can edit it instead",
            );
          else this.notify.error('An unknown error occurred');
        },
      });
  }

  onNavigateBack(): void {
    this.store.dispatch(new Navigate(['/connections']));
  }
}
