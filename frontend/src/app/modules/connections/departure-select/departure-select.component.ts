import {
  Component,
  EventEmitter,
  Input,
  OnDestroy,
  Output,
  TrackByFunction,
} from '@angular/core';
import { BahnPlace, BahnService } from '../../../api';
import { RelativeTime } from '../relative-time/relative-time';
import { FormControl } from '@angular/forms';
import {
  BehaviorSubject,
  combineLatest,
  EMPTY,
  map,
  shareReplay,
  startWith,
  switchMap,
} from 'rxjs';

@Component({
  selector: 'app-departure-select',
  templateUrl: './departure-select.component.html',
  styleUrls: ['./departure-select.component.scss'],
})
export class DepartureSelectComponent implements OnDestroy {
  @Input() set from(station: BahnPlace | undefined) {
    this.from$.next(station);
  }

  get from(): BahnPlace | undefined {
    return this.from$.value;
  }

  @Input() set to(station: BahnPlace | undefined) {
    this.to$.next(station);
  }

  get to(): BahnPlace | undefined {
    return this.to$.value;
  }

  private readonly from$ = new BehaviorSubject<BahnPlace | undefined>(
    undefined,
  );
  private readonly to$ = new BehaviorSubject<BahnPlace | undefined>(undefined);

  targetDeparture = new FormControl(RelativeTime.now(), { nonNullable: true });
  @Input() selectedDepartures: ReadonlyArray<RelativeTime> = [];
  @Output() selectedDeparturesChange = new EventEmitter<
    ReadonlyArray<RelativeTime>
  >();

  foundDepartures$ = combineLatest([
    this.from$,
    this.to$,
    this.targetDeparture.valueChanges.pipe(
      startWith(this.targetDeparture.value),
    ),
  ]).pipe(
    switchMap(([from, to, departure]) => {
      if (!from || !to) return EMPTY;

      return this.bahn
        .bahnConnectionsGet(departure.toIso(), from.stationId, to.stationId)
        .pipe(startWith(undefined));
    }),
    map((response) =>
      response?.connections.map((c) =>
        RelativeTime.fromIso(c.departure.scheduledTime),
      ),
    ),
    shareReplay(1),
  );

  constructor(private readonly bahn: BahnService) {}

  ngOnDestroy(): void {
    this.selectedDeparturesChange.complete();
  }

  onToggle(departure: RelativeTime): void {
    const selectedDepartureIndex = this.selectedDepartures.findIndex(
      (d) => d.str === departure.str,
    );

    if (selectedDepartureIndex === -1) {
      this.selectedDepartures = [departure, ...this.selectedDepartures];
    } else {
      this.selectedDepartures = this.selectedDepartures.filter(
        (_, index) => index !== selectedDepartureIndex,
      );
    }

    this.selectedDeparturesChange.emit(this.selectedDepartures);
  }

  trackByDeparture: TrackByFunction<RelativeTime> = (_, { timestamp }) =>
    timestamp;
}
