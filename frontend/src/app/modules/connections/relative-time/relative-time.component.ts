import { Component, forwardRef, Input, OnDestroy, OnInit } from '@angular/core';
import { ControlValueAccessor, FormControl, NG_VALUE_ACCESSOR, ReactiveFormsModule } from '@angular/forms';
import { RelativeTime } from './relative-time';
import { Subject, takeUntil } from 'rxjs';
import { NextRelativeTimePipe } from '../../shared/pipes/next-relative-time.pipe';
import { AsyncPipe } from '@angular/common';

@Component({
    selector: 'app-relative-time',
    templateUrl: './relative-time.component.html',
    styleUrls: ['./relative-time.component.scss'],
    providers: [
        {
            provide: NG_VALUE_ACCESSOR,
            multi: true,
            useExisting: forwardRef(() => RelativeTimeComponent),
        },
    ],
    standalone: true,
    imports: [
        ReactiveFormsModule,
        AsyncPipe,
        NextRelativeTimePipe,
    ],
})
export class RelativeTimeComponent
  implements ControlValueAccessor, OnInit, OnDestroy
{
  @Input() inputId?: string;

  time = RelativeTime.now();

  readonly inputControl = new FormControl(this.time.str, {
    nonNullable: true,
  });

  private onTouched?: () => void;
  private onChange?: (time: RelativeTime) => void;
  private readonly destroy$ = new Subject<void>();

  ngOnInit(): void {
    this.inputControl.valueChanges
      .pipe(takeUntil(this.destroy$))
      .subscribe((time) => {
        this.time = RelativeTime.fromTimeInput(time);
        if (this.onChange) this.onChange(this.time);
      });
  }

  ngOnDestroy(): void {
    this.destroy$.next();
    this.destroy$.complete();
  }

  writeValue(time: RelativeTime): void {
    this.time = time;
    this.inputControl.setValue(time.str);
  }

  setDisabledState(isDisabled: boolean): void {
    if (isDisabled) this.inputControl.disable();
    else this.inputControl.enable();
  }

  registerOnChange(fn: (time: RelativeTime) => void) {
    this.onChange = fn;
  }

  registerOnTouched(fn: () => void) {
    this.onTouched = fn;
  }

  onBlur(): void {
    if (this.onTouched) this.onTouched();
  }
}
