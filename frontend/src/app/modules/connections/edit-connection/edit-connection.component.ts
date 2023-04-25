import { Component } from '@angular/core';
import { RelativeTime } from '../../shared/relative-time/relative-time';
import { FormBuilder, Validators } from '@angular/forms';
import { BahnPlace } from '../../../api';

@Component({
  selector: 'app-edit-connection',
  templateUrl: './edit-connection.component.html',
  styleUrls: ['./edit-connection.component.scss'],
})
export class EditConnectionComponent {
  // readonly form = this.fb.nonNullable.group({
  //   from: [null as BahnPlace | null, Validators.required],
  //   to: [null as BahnPlace | null, Validators.required],
  // });

  readonly form = this.fb.nonNullable.group({
    from: [
      {
        id: '38Dfr1pLBwweELSuarsr9ouX8',
        label: 'Uetze',
        name: 'Dedenhausen',
        stationId: '8001392',
      } as BahnPlace | null,
      Validators.required,
    ],
    to: [
      {
        id: '38Dfr1pL2y4tedsPCBKNPRSYi',
        label: 'Mitte, Hannover',
        name: 'Hannover Hbf',
        stationId: '8000152',
      } as BahnPlace | null,
      Validators.required,
    ],
  });

  selectedDepartures: Array<RelativeTime> = [
    RelativeTime.now(),
    RelativeTime.now(),
  ];

  constructor(private readonly fb: FormBuilder) {}
}
