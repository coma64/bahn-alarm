import {
  Component,
  ElementRef,
  forwardRef,
  Input,
  OnDestroy,
  TrackByFunction,
  ViewChild,
} from '@angular/core';
import {
  ControlValueAccessor,
  FormControl,
  NG_VALUE_ACCESSOR,
} from '@angular/forms';
import { CdkPortal, TemplatePortal } from '@angular/cdk/portal';
import { Overlay } from '@angular/cdk/overlay';
import {
  catchError,
  debounceTime,
  distinctUntilChanged,
  EMPTY,
  of,
  shareReplay,
  startWith,
  Subject,
  switchMap,
  take,
  takeUntil,
} from 'rxjs';
import { BahnService, BahnStation } from '../../../api';
import { NotifyService } from '../../shared/services/notify.service';

@Component({
  selector: 'app-station-search',
  templateUrl: './station-search.component.html',
  styleUrls: ['./station-search.component.scss'],
  providers: [
    {
      provide: NG_VALUE_ACCESSOR,
      multi: true,
      useExisting: forwardRef(() => StationSearchComponent),
    },
  ],
})
export class StationSearchComponent implements ControlValueAccessor, OnDestroy {
  @Input() inputId?: string;

  @ViewChild(CdkPortal) readonly suggestionsPortal?: TemplatePortal;
  @ViewChild('suggestionsElement')
  readonly suggestionsElement?: ElementRef<HTMLDivElement>;

  readonly inputControl = new FormControl('', { nonNullable: true });

  readonly suggestions$ = this.inputControl.valueChanges.pipe(
    distinctUntilChanged(),
    debounceTime(200),
    switchMap((stationName) => {
      if (stationName.length < 3) return EMPTY;

      // Start with undefined to show the loading spinner again
      return this.bahn.bahnPlacesGet(stationName).pipe(startWith(undefined));
    }),
    catchError(() => {
      this.notify.error('An error occurred while loading the stations');
      return of(undefined);
    }),
    shareReplay(1),
  );

  @ViewChild('inputElement')
  private readonly inputElement?: ElementRef<HTMLInputElement>;

  private readonly destroy$ = new Subject<void>();
  private onTouched?: () => void;
  private onChange?: (station: BahnStation) => void;

  constructor(
    private readonly overlay: Overlay,
    private readonly bahn: BahnService,
    private readonly notify: NotifyService,
  ) {}

  ngOnDestroy(): void {
    this.destroy$.next();
    this.destroy$.complete();
  }

  writeValue(station?: BahnStation) {
    this.inputControl.setValue(station?.name ?? '');
  }

  registerOnChange(fn: (station: BahnStation) => void) {
    this.onChange = fn;
  }

  registerOnTouched(fn: () => void) {
    this.onTouched = fn;
  }

  setDisabledState(isDisabled: boolean) {
    if (isDisabled) this.inputControl.disable();
    else this.inputControl.enable();

    this.hide();
  }

  onBlur(event: FocusEvent): void {
    if (
      this.suggestionsElement?.nativeElement.contains(
        event.relatedTarget as Node,
      )
    )
      return;

    this.hide();
    if (this.onTouched) this.onTouched();
  }

  onSelect(station: BahnStation): void {
    this.inputControl.setValue(station.name, { emitEvent: false });
    if (this.onChange) this.onChange(station);
    this.hide();
  }

  onEnter(): void {
    this.suggestions$
      .pipe(take(1), takeUntil(this.destroy$))
      .subscribe((suggestions) => {
        if (suggestions?.places.length) {
          this.onSelect(suggestions.places[0]);
        }
      });
  }

  show(): void {
    if (
      !this.inputElement ||
      !this.suggestionsPortal ||
      this.suggestionsPortal.isAttached
    )
      return;

    const positionStrategy = this.overlay
      .position()
      .flexibleConnectedTo(this.inputElement.nativeElement)
      .withPositions([
        {
          originX: 'center',
          overlayX: 'center',
          originY: 'bottom',
          overlayY: 'top',
        },
      ]);

    const overlay = this.overlay.create({
      positionStrategy,
      scrollStrategy: this.overlay.scrollStrategies.reposition(),
    });

    this.suggestionsPortal?.attach(overlay);
  }

  hide(): void {
    if (this.suggestionsPortal?.isAttached) this.suggestionsPortal.detach();
  }

  trackByPlace: TrackByFunction<BahnStation> = (_, { id }) => id;
}
