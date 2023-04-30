import {
  Component,
  ElementRef,
  EventEmitter,
  HostBinding,
  HostListener,
  Input,
  OnDestroy,
  Output,
  TrackByFunction,
  ViewChild,
} from '@angular/core';
import { CdkPortal, TemplatePortal } from '@angular/cdk/portal';
import { Overlay } from '@angular/cdk/overlay';

export type Option<T> = {
  label: string;
  value: T;
};

@Component({
  // eslint-disable-next-line @angular-eslint/component-selector
  selector: 'button[app-dropdown]',
  templateUrl: './dropdown.component.html',
  styleUrls: ['./dropdown.component.scss'],
})
export class DropdownComponent<T> implements OnDestroy {
  @Input() options: ReadonlyArray<Option<T>> = [];

  @Input() selected?: Option<T>;
  @Output() readonly selectedChange = new EventEmitter<Option<T>>();

  @ViewChild(CdkPortal) private readonly portal?: TemplatePortal;

  get isOpen(): boolean {
    return Boolean(this.portal?.isAttached);
  }

  @HostBinding('disabled')
  get isDisabled(): boolean {
    return !this.options.length;
  }

  constructor(
    private readonly overlay: Overlay,
    readonly element: ElementRef<HTMLButtonElement>,
  ) {}

  ngOnDestroy() {
    this.selectedChange.complete();
  }

  onSelect(value: Option<T>): void {
    this.selected = value;
    this.selectedChange.next(value);
    this.hide();
  }

  hide(): void {
    if (this.isOpen) this.portal?.detach();
  }

  trackByOption: TrackByFunction<Option<T>> = (_, { value }) => value;

  @HostListener('click')
  private onClick(): void {
    if (this.isDisabled) return;

    if (this.isOpen) this.hide();
    else this.show();
  }

  @HostListener('window:click', ['$event'])
  private onWindowClick(event: MouseEvent): void {
    if (!this.element.nativeElement.contains(event.target as Node)) {
      this.hide();
    }
  }

  private show(): void {
    if (!this.portal) return;

    const positionStrategy = this.overlay
      .position()
      .flexibleConnectedTo(this.element.nativeElement)
      .withPositions([
        {
          originX: 'center',
          overlayX: 'center',
          originY: 'bottom',
          overlayY: 'top',
        },
        {
          originX: 'center',
          overlayX: 'center',
          originY: 'top',
          overlayY: 'bottom',
        },
      ]);

    const scrollStrategy = this.overlay.scrollStrategies.reposition();
    const overlay = this.overlay.create({
      positionStrategy,
      scrollStrategy,
    });

    this.portal.attach(overlay);
  }
}
