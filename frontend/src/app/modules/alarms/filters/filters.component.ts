import { Component } from '@angular/core';
import {
  DropdownComponent,
  Option,
} from '../../shared/dropdown/dropdown.component';
import { Urgency } from '../../../api';
import { Store } from '@ngxs/store';
import { AlarmsActions } from '../../../state/alarms.actions';
import { IconsModule } from '../../icons/icons.module';
import { NotifyService } from '../../shared/services/notify.service';
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';

@Component({
  selector: 'app-filters',
  templateUrl: './filters.component.html',
  styleUrls: ['./filters.component.scss'],
  standalone: true,
  imports: [IconsModule, DropdownComponent],
})
export class FiltersComponent {
  readonly urgencyOptions: Array<Option<Urgency | undefined>> = [
    { label: 'All', value: undefined },
    { label: 'Info', value: Urgency.Info },
    { label: 'Warning', value: Urgency.Warn },
    { label: 'Error', value: Urgency.Error },
  ];

  selectedUrgency: Option<Urgency | undefined> = this.urgencyOptions[0];

  private readonly untilDestroyed = takeUntilDestroyed();

  constructor(
    private readonly store: Store,
    private readonly notify: NotifyService,
  ) {}

  onUrgencyChange({ value }: Option<Urgency | undefined>): void {
    this.store.dispatch(new AlarmsActions.FilterByUrgency(value));
  }

  onRefresh(): void {
    this.store.dispatch(AlarmsActions.Fetch);
  }

  onDelete(): void {
    this.notify
      .confirm('Are you sure you want to delete all alarms?')
      .pipe(this.untilDestroyed)
      .subscribe(
        (confirmed) =>
          confirmed && this.store.dispatch(AlarmsActions.DeleteAll),
      );
  }
}
