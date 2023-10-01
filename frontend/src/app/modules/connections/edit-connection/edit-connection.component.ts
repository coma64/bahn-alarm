import { Component, OnDestroy, OnInit } from '@angular/core';
import { RelativeTime } from '../relative-time/relative-time';
import { FormBuilder, Validators, ReactiveFormsModule } from '@angular/forms';
import { BahnStation, TrackingService } from '../../../api';
import { Store } from '@ngxs/store';
import { Navigate } from '@ngxs/router-plugin';
import { HttpErrorResponse } from '@angular/common/http';
import { EMPTY, finalize, map, Subject, switchMap, takeUntil, tap } from 'rxjs';
import { NotifyService } from '../../shared/services/notify.service';
import { Connections } from '../../../state/connections.actions';
import { ActivatedRoute } from '@angular/router';
import { SpinnerComponent } from '../../shared/components/spinner/spinner.component';
import { BannerComponent } from '../../shared/components/banner/banner.component';
import { DepartureSelectComponent } from '../departure-select/departure-select.component';
import { StationSearchComponent } from '../station-search/station-search.component';
import { FeatherModule } from 'angular-feather';
import { NgIf } from '@angular/common';

@Component({
    selector: 'app-edit-connection',
    templateUrl: './edit-connection.component.html',
    styleUrls: ['./edit-connection.component.scss'],
    standalone: true,
    imports: [
        NgIf,
        FeatherModule,
        ReactiveFormsModule,
        StationSearchComponent,
        DepartureSelectComponent,
        BannerComponent,
        SpinnerComponent,
    ],
})
export class EditConnectionComponent implements OnInit, OnDestroy {
  readonly stationForm = this.fb.nonNullable.group({
    from: [undefined as BahnStation | undefined, Validators.required],
    to: [undefined as BahnStation | undefined, Validators.required],
  });

  selectedDepartures: readonly RelativeTime[] = [];
  hasTriedSubmitting = false;
  isEditing = false;
  isLoading = true;

  get from(): BahnStation | undefined {
    return this.stationForm.value.from;
  }

  get to(): BahnStation | undefined {
    return this.stationForm.value.to;
  }

  private connectionId?: number;
  private readonly destroy$ = new Subject<void>();

  constructor(
    private readonly fb: FormBuilder,
    private readonly tracking: TrackingService,
    private readonly notify: NotifyService,
    private readonly activatedRoute: ActivatedRoute,
    protected readonly store: Store,
  ) {}

  ngOnInit(): void {
    this.activatedRoute.params
      .pipe(
        map(({ connectionId }) =>
          connectionId ? Number(connectionId) : undefined,
        ),
        tap((id) => {
          this.isLoading = this.isEditing = !!id;
          this.connectionId = id;
          if (id !== undefined) {
            this.stationForm.disable();
          } else {
            this.stationForm.enable();
          }
        }),
        switchMap((id) =>
          id
            ? this.tracking
                .trackingConnectionsIdGet(id)
                .pipe(finalize(() => (this.isLoading = false)))
            : EMPTY,
        ),
        takeUntil(this.destroy$),
      )
      .subscribe(
        ({ from, to, departures }) => {
          this.stationForm.setValue({ from, to });
          this.selectedDepartures = departures.map((d) =>
            RelativeTime.fromIso(d.departure),
          );
        },
        () =>
          this.notify.error('An error occurred while loading the connection'),
      );

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
    const apiConnection = {
      from: this.from,
      to: this.to,
      departures: this.selectedDepartures.map((d) => ({
        departure: d.toIsoZeroBased(),
      })),
    };

    if (this.isEditing) {
      if (!this.connectionId) return;
      this.tracking
        .trackingConnectionsIdPut(this.connectionId, apiConnection)
        .subscribe({
          next: (connection) =>
            this.store.dispatch([
              new Connections.Updated(connection),
              new Navigate(['/connections']),
            ]),
          error: () => this.notify.error('An unknown error occurred'),
        });
    } else {
      this.tracking.trackingConnectionsPost(apiConnection).subscribe({
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
  }

  onNavigateBack(): void {
    this.store.dispatch(new Navigate(['/connections']));
  }
}
