import {
  Component,
  EventEmitter,
  Input,
  OnDestroy,
  Output,
} from '@angular/core';
import { BahnPlace, BahnService } from '../../../api';
import { RelativeTime } from '../../shared/relative-time/relative-time';
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
  @Input() set from(station: BahnPlace | undefined | null) {
    this.from$.next(station ?? undefined);
  }

  get from(): BahnPlace | undefined {
    return this.from$.value;
  }

  @Input() set to(station: BahnPlace | undefined | null) {
    this.to$.next(station ?? undefined);
  }

  get to(): BahnPlace | undefined {
    return this.to$.value;
  }

  private readonly from$ = new BehaviorSubject<BahnPlace | undefined>(
    undefined,
  );
  private readonly to$ = new BehaviorSubject<BahnPlace | undefined>(undefined);

  targetDeparture = new FormControl(RelativeTime.now(), { nonNullable: true });
  @Input() selectedDepartures: Array<RelativeTime> = [];
  @Output() selectedDeparturesChange = new EventEmitter<Array<RelativeTime>>();

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
      this.selectedDepartures.splice(selectedDepartureIndex, 1);
      this.selectedDepartures = [...this.selectedDepartures];
    }

    this.selectedDeparturesChange.emit(this.selectedDepartures);
  }
}
