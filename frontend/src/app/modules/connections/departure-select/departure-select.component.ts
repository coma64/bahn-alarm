import {
  Component,
  EventEmitter,
  Input,
  OnDestroy,
  Output,
  TrackByFunction,
} from '@angular/core';
import { BahnService, BahnStation } from '../../../api';
import { RelativeTime } from '../relative-time/relative-time';
import { FormControl } from '@angular/forms';
import {
  BehaviorSubject,
  combineLatest,
  EMPTY,
  map,
  Observable,
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
  get from(): BahnStation | undefined {
    return this.from$.value;
  }

  @Input() set from(station: BahnStation | undefined) {
    this.from$.next(station);
  }

  get to(): BahnStation | undefined {
    return this.to$.value;
  }

  @Input() set to(station: BahnStation | undefined) {
    this.to$.next(station);
  }

  @Input() selectedDepartures: readonly RelativeTime[] = [];

  readonly targetDeparture = new FormControl(RelativeTime.now(), {
    nonNullable: true,
  });

  protected readonly foundDepartures$: Observable<RelativeTime[] | undefined>;

  @Output() private readonly selectedDeparturesChange = new EventEmitter<
    readonly RelativeTime[]
  >();

  private readonly from$ = new BehaviorSubject<BahnStation | undefined>(
    undefined,
  );

  private readonly to$ = new BehaviorSubject<BahnStation | undefined>(
    undefined,
  );

  constructor(private readonly bahn: BahnService) {
    this.foundDepartures$ = combineLatest([
      this.from$,
      this.to$,
      this.targetDeparture.valueChanges.pipe(
        startWith(this.targetDeparture.value),
      ),
    ]).pipe(
      switchMap(([from, to, departure]) => {
        if (!from || !to) return EMPTY;

        return this.bahn
          .bahnConnectionsGet(departure.toIso(), from.id, to.id)
          .pipe(startWith(undefined));
      }),
      map((response) =>
        response?.connections.map((c) =>
          RelativeTime.fromIso(c.departure.scheduledTime),
        ),
      ),
      shareReplay(1),
    );
  }

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
