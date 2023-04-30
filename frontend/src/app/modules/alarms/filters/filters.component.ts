import { Component } from '@angular/core';
import { Option } from '../../shared/dropdown/dropdown.component';
import { Urgency } from '../../../api';
import { Store } from '@ngxs/store';
import { AlarmsActions } from '../../../state/alarms.actions';

@Component({
  selector: 'app-filters',
  templateUrl: './filters.component.html',
  styleUrls: ['./filters.component.scss'],
})
export class FiltersComponent {
  readonly urgencyOptions: Array<Option<Urgency | undefined>> = [
    { label: 'All', value: undefined },
    { label: 'Info', value: Urgency.Info },
    { label: 'Warning', value: Urgency.Warn },
    { label: 'Error', value: Urgency.Error },
  ];

  selectedUrgency: Option<Urgency | undefined> = this.urgencyOptions[0];

  constructor(private readonly store: Store) {}

  onUrgencyChange({ value }: Option<Urgency | undefined>): void {
    this.store.dispatch(new AlarmsActions.FilterByUrgency(value));
  }

  onRefresh(): void {
    this.store.dispatch(AlarmsActions.Fetch);
  }
}
