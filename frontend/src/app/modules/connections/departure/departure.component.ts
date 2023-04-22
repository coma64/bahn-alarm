import { Component, Input } from '@angular/core';
import { TrackedDeparture } from '../../../api';

@Component({
  selector: 'app-departure',
  templateUrl: './departure.component.html',
  styleUrls: ['./departure.component.scss'],
})
export class DepartureComponent {
  @Input() departure?: TrackedDeparture;
}
