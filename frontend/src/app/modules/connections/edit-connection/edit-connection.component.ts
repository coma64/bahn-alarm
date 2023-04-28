import { Component, OnDestroy, OnInit } from '@angular/core';
import { RelativeTime } from '../relative-time/relative-time';
import { FormBuilder, Validators } from '@angular/forms';
import { BahnPlace, TrackingService } from '../../../api';
import { Store } from '@ngxs/store';
import { Connections } from '../../../state/connections.actions';
import { Navigate } from '@ngxs/router-plugin';
import { HttpErrorResponse } from '@angular/common/http';
import { Subject, takeUntil } from 'rxjs';

@Component({
  selector: 'app-edit-connection',
  templateUrl: './edit-connection.component.html',
  styleUrls: ['./edit-connection.component.scss'],
})
export class EditConnectionComponent implements OnInit, OnDestroy {
  readonly stationForm = this.fb.group({
    from: [null as BahnPlace | null, Validators.required],
    to: [null as BahnPlace | null, Validators.required],
  });

  selectedDepartures: ReadonlyArray<RelativeTime> = [];
  hasTriedSubmitting = false;

  get from(): BahnPlace | undefined {
    return this.stationForm.value.from ?? undefined;
  }

  get to(): BahnPlace | undefined {
    return this.stationForm.value.to ?? undefined;
  }

  private readonly destroy$ = new Subject<void>();

  constructor(
    private readonly fb: FormBuilder,
    private readonly tracking: TrackingService,
    readonly store: Store,
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
      confirm(
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
            alert(
              "You're already tracking this connection. You can edit it instead",
            );
          else alert('An unknown error occurred');
        },
      });
  }

  onNavigateBack(): void {
    this.store.dispatch(new Navigate(['/connections']));
  }
}
