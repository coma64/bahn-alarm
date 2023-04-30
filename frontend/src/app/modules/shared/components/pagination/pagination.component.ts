import {
  Component,
  EventEmitter,
  Input,
  OnDestroy,
  OnInit,
  Output,
} from '@angular/core';
import { FormControl } from '@angular/forms';
import { debounceTime, Subject, takeUntil } from 'rxjs';

@Component({
  selector: 'app-pagination',
  templateUrl: './pagination.component.html',
  styleUrls: ['./pagination.component.scss'],
})
export class PaginationComponent implements OnInit, OnDestroy {
  get page(): number {
    return this.control.value - 1;
  }

  @Input() set page(value: number) {
    this.control.setValue(value + 1);
    if (value === this.maxPage) this.control.disable();
    else this.control.enable();
  }

  get shownPage(): number {
    return this.control.value;
  }

  get maxPage(): number {
    return this._maxPage;
  }

  @Input() set maxPage(value: number) {
    this._maxPage = value;
    if (this.page === value) this.control.disable();
    else this.control.enable();
  }

  @Output() protected readonly pageChange = new EventEmitter<number>();

  protected readonly control = new FormControl(1, { nonNullable: true });
  private readonly destroy$ = new Subject<void>();
  private _maxPage = 0;

  ngOnInit(): void {
    this.control.valueChanges
      .pipe(debounceTime(200), takeUntil(this.destroy$))
      .subscribe((newPage) => this.pageChange.emit(newPage - 1));
  }

  ngOnDestroy(): void {
    this.destroy$.next();
    this.destroy$.complete();
  }

  increment(): void {
    this.page += 1;
  }

  decrement(): void {
    this.page -= 1;
  }
}
