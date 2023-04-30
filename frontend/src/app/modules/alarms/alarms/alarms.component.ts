import { Component, OnInit } from '@angular/core';
import { Select, Store } from '@ngxs/store';
import { AlarmsActions } from '../../../state/alarms.actions';
import { AlarmsState, AlarmsStateModel } from '../../../state/alarms.state';
import { Observable } from 'rxjs';

@Component({
  selector: 'app-alarms',
  templateUrl: './alarms.component.html',
  styleUrls: ['./alarms.component.scss'],
})
export class AlarmsComponent implements OnInit {
  @Select(AlarmsState)
  protected readonly alarms$!: Observable<AlarmsStateModel>;

  constructor(private readonly store: Store) {}

  ngOnInit() {
    if (!this.store.selectSnapshot<AlarmsStateModel>(AlarmsState).items)
      this.store.dispatch(AlarmsActions.Fetch);
  }

  changePage(newPage: number): void {
    this.store.dispatch(new AlarmsActions.SetPage(newPage));
  }
}
